package handlers

import (
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

		// search := buildSearchQuery()

		// http.Error(w, err.Error(), http.StatusInternalServerError)
		// json.NewEncoder(w).Encode(customer)
	}
}

func buildSearchQuery() *searchQuery {
	return &searchQuery{
		query: queries.BaseRoomSearch,
	}
}

func (q *searchQuery) addFilter(filter string) {
	if q.filterCount == 0 {
		q.query += fmt.Sprintf(" WHERE %s", filter)
		return
	}

	q.query += fmt.Sprintf(" AND %s", filter)
	q.filterCount += 1
}

func (q *searchQuery) withCapacityFilter(value int) *searchQuery {
	q.addFilter(fmt.Sprintf("capacity >= $%v", q.filterCount))

	q.capacity.applied = true
	q.capacity.val = value

	return q
}

func (q *searchQuery) withChainNameFilter(value string) *searchQuery {
	q.addFilter(fmt.Sprintf("HotelChains.chain_name LIKE '%%$%v%%'", q.filterCount))

	q.chainName.applied = true
	q.chainName.val = value

	return q
}
