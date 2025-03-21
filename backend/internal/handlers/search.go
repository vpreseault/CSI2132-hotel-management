package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vpreseault/csi2132-project/backend/internal"
	"github.com/vpreseault/csi2132-project/backend/internal/queries"
)

type filter[T string | int | float32] struct {
	applied bool
	val     T
}

type searchQuery struct {
	query       string
	filterCount int
	capacity    filter[int]
	price       filter[int] // TODO: allow for price range
	chainName   filter[string]
}

// - [ ] start/end date
// - [x] room capacity
// - [ ] area/city
// - [x] hotel chain
// - [ ] hotel category
// - [ ] total number of rooms in hotel
// - [x] room price

func RoomSearchHandler(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var searchParams struct {
			Capacity  *int    `json:"capacity,omitempty"`
			ChainName *string `json:"chain_name,omitempty"`
			Price     *int    `json:"price,omitempty"`
		}

		if err := json.NewDecoder(r.Body).Decode(&searchParams); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		search := buildSearchQuery()

		if searchParams.Capacity != nil {
			search.withCapacityFilter(*searchParams.Capacity)
		}

		if searchParams.ChainName != nil {
			search.withChainNameFilter(*searchParams.ChainName)
		}

		rooms, err := search.executeQuery(ctx)
		if err != nil {
			fmt.Println(search.query)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(rooms)
	}
}

func buildSearchQuery() *searchQuery {
	return &searchQuery{
		query:       queries.BaseRoomSearch,
		filterCount: 1,
	}
}

func (q *searchQuery) executeQuery(ctx *internal.AppContext) ([]internal.SearchResult, error) {
	var searchResults []internal.SearchResult
	var args []interface{}

	if q.capacity.applied {
		args = append(args, q.capacity.val)
	}
	if q.chainName.applied {
		args = append(args, q.chainName.val)
	}

	fmt.Println(args...)
	rows, err := ctx.DB.Query(q.query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var searchResult internal.SearchResult
		if err := rows.Scan(
			&searchResult.ID,
			&searchResult.HotelID,
			&searchResult.RoomNumber,
			&searchResult.Capacity,
			&searchResult.Price,
			&searchResult.ViewType,
			&searchResult.Extendable,
			&searchResult.Damaged,
			&searchResult.ChainName,
			&searchResult.Category,
			&searchResult.Address,
			&searchResult.TotalRooms,
		); err != nil {
			return nil, err
		}
		searchResults = append(searchResults, searchResult)
	}

	return searchResults, rows.Err()
}

func (q *searchQuery) addFilter(filter string) {
	if q.filterCount == 1 {
		q.query += fmt.Sprintf(" WHERE %s", filter)
	} else {
		q.query += fmt.Sprintf(" AND %s", filter)
	}
	q.filterCount++
}

func (q *searchQuery) withCapacityFilter(value int) *searchQuery {
	q.addFilter(fmt.Sprintf("capacity >= $%v", q.filterCount))

	q.capacity.applied = true
	q.capacity.val = value

	return q
}

func (q *searchQuery) withChainNameFilter(value string) *searchQuery {
	q.addFilter(fmt.Sprintf("chain_name LIKE CONCAT('%%', CAST($%v AS TEXT), '%%')", q.filterCount))
	q.chainName.applied = true
	q.chainName.val = value
	return q
}
