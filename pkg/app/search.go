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
	var x, y string
	urlParams := r.URL.Query()
	if p := urlParams.Get("price"); IsValid(p) {
		x = p
	}
	if p := urlParams.Get("date"); IsValid(p) {
		y = p
	}
	t := ds.List(x, y)
	msg.ReturnJSON(w, t)
}

func IsValid(s string) bool {
	switch s {
	case "ACS", "DESC":
		return true
	}
	return false
}
