package app

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rodkevich/tbd/internal/msg"
)

// View ...
func (a Application) View(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	key := r.FormValue("fields")
	if key != "" {
		q := ds.ViewByID(id)
		msg.ReturnJSON(w, q)
		return
	}
	msg.ReturnServerError(w, errors.New("not yet implemented"))
}
