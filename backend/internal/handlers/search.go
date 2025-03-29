package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vpreseault/csi2132-project/backend/internal"
	"github.com/vpreseault/csi2132-project/backend/internal/queries"
)

type searchQuery struct {
	query       string
	filterCount int
	args        []interface{}
}

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
	sq := &searchQuery{
		query:       queries.BaseRoomSearch,
		filterCount: 1,
	}

	if params.Capacity != nil {
		sq.addFilter(fmt.Sprintf("capacity >= $%v", sq.filterCount))
		sq.args = append(sq.args, *params.Capacity)
	}

	if params.ChainName != nil {
		sq.addFilter(fmt.Sprintf("chain_name LIKE CONCAT('%%', CAST($%v AS TEXT), '%%')", sq.filterCount))
		sq.args = append(sq.args, *params.ChainName)
	}

	if params.MaxPrice != nil {
		sq.addFilter(fmt.Sprintf("price <= $%v", sq.filterCount))
		sq.args = append(sq.args, *params.MaxPrice)
	}

	if params.Area != nil {
		sq.addFilter(fmt.Sprintf("address LIKE CONCAT('%%', CAST($%v AS TEXT), '%%')", sq.filterCount))
		sq.args = append(sq.args, *params.Area)
	}

	if params.Category != nil {
		sq.addFilter(fmt.Sprintf("category = $%v", sq.filterCount))
		sq.args = append(sq.args, *params.Category)
	}

	if params.TotalRooms != nil {
		sq.addFilter(fmt.Sprintf("total_rooms = $%v", sq.filterCount))
		sq.args = append(sq.args, *params.TotalRooms)
	}

	if params.StartDate != nil && params.EndDate != nil {
		sq.addFilter(fmt.Sprintf("(is_available = true OR ((renting_start NOT BETWEEN $%v AND $%v) AND (renting_end NOT BETWEEN $%v AND $%v)))",
			sq.filterCount, sq.filterCount+1, sq.filterCount, sq.filterCount+1))
		sq.filterCount++
		sq.args = append(sq.args, *params.StartDate)
		sq.args = append(sq.args, *params.EndDate)
	}

	sq.query += " LIMIT 6"

	return sq
}

func (q *searchQuery) executeQuery(ctx *internal.AppContext) ([]internal.SearchResult, error) {
	var searchResults []internal.SearchResult

	rows, err := ctx.DB.Query(q.query, q.args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var searchResult internal.SearchResult
		if err := rows.Scan(
			&searchResult.ID,
			&searchResult.HotelID,
			&searchResult.HotelName,
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
