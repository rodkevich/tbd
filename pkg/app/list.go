package app

import (
	"net/http"

	"github.com/rodkevich/tbd/internal/msg"
)

// SearchConfig ...
type SearchConfig struct {
}

// List ...
func (a Application) List(w http.ResponseWriter, r *http.Request) {
	var PriceSort, DateSort string
	urlParams := r.URL.Query()
	if p := urlParams.Get("price"); IsValid(p) {
		PriceSort = p
	}
	if p := urlParams.Get("date"); IsValid(p) {
		DateSort = p
	}
	tickets := ds.List(PriceSort, DateSort)
	msg.ReturnJSON(w, tickets)
}

func IsValid(s string) bool {
	switch s {
	case "ASC", "DESC":
		return true
	}
	return false
}
