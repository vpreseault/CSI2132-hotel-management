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
	maxPrice    filter[int] // TODO: allow for price range
	category    filter[int] // TODO: allow for category range?
	chainName   filter[string]
	area        filter[string]
	totalRooms  filter[int]
	startDate   filter[string]
	endDate     filter[string]
}

// - [x] start/end date
// - [x] room capacity
// - [x] area/city
// - [x] hotel chain
// - [x] hotel category
// - [x] total number of rooms in hotel
// - [x] room price

type searchParams struct {
	Capacity   *int    `json:"capacity,omitempty"`
	ChainName  *string `json:"chain_name,omitempty"`
	Area       *string `json:"area,omitempty"`
	Category   *int    `json:"category,omitempty"`
	MaxPrice   *int    `json:"max_price,omitempty"`
	TotalRooms *int    `json:"total_rooms,omitempty"`
	StartDate  *string `json:"start_date,omitempty"`
	EndDate    *string `json:"end_date,omitempty"`
}

func RoomSearchHandler(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var params searchParams

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		query := buildSearchQuery(params)

		rooms, err := query.executeQuery(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(rooms)
	}
}

func buildSearchQuery(params searchParams) *searchQuery {
	search := &searchQuery{
		query:       queries.BaseRoomSearch,
		filterCount: 1,
	}

	if params.Capacity != nil {
		search.withCapacityFilter(*params.Capacity)
	}

	if params.ChainName != nil {
		search.withChainNameFilter(*params.ChainName)
	}

	if params.MaxPrice != nil {
		search.withMaxPriceFilter(*params.MaxPrice)
	}

	if params.Area != nil {
		search.withAreaFilter(*params.Area)
	}

	if params.Category != nil {
		search.withCategoryFilter(*params.Category)
	}

	if params.TotalRooms != nil {
		search.withTotalRoomsFilter(*params.TotalRooms)
	}

	if params.StartDate != nil && params.EndDate != nil {
		search.withDateRangeFilter(*params.StartDate, *params.EndDate)
	}

	return search
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
	if q.maxPrice.applied {
		args = append(args, q.maxPrice.val)
	}
	if q.area.applied {
		args = append(args, q.area.val)
	}
	if q.category.applied {
		args = append(args, q.category.val)
	}
	if q.totalRooms.applied {
		args = append(args, q.totalRooms.val)
	}
	if q.startDate.applied && q.endDate.applied {
		args = append(args, q.startDate.val)
		args = append(args, q.endDate.val)
	}

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

func (q *searchQuery) withAreaFilter(value string) *searchQuery {
	q.addFilter(fmt.Sprintf("address LIKE CONCAT('%%', CAST($%v AS TEXT), '%%')", q.filterCount))
	q.area.applied = true
	q.area.val = value
	return q
}

func (q *searchQuery) withMaxPriceFilter(value int) *searchQuery {
	q.addFilter(fmt.Sprintf("price <= $%v", q.filterCount))
	q.maxPrice.applied = true
	q.maxPrice.val = value
	return q
}

func (q *searchQuery) withCategoryFilter(value int) *searchQuery {
	q.addFilter(fmt.Sprintf("category = $%v", q.filterCount))
	q.category.applied = true
	q.category.val = value
	return q
}

func (q *searchQuery) withTotalRoomsFilter(value int) *searchQuery {
	q.addFilter(fmt.Sprintf("total_rooms = $%v", q.filterCount))
	q.totalRooms.applied = true
	q.totalRooms.val = value
	return q
}

func (q *searchQuery) withDateRangeFilter(start, end string) *searchQuery {
	q.addFilter(fmt.Sprintf("(is_available = true OR ((renting_start NOT BETWEEN $%v AND $%v) AND (renting_end NOT BETWEEN $%v AND $%v)))",
		q.filterCount, q.filterCount+1, q.filterCount, q.filterCount+1))
	q.filterCount += 1
	q.startDate.applied = true
	q.startDate.val = start
	q.endDate.applied = true
	q.endDate.val = end
	return q
}
