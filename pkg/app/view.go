package app

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/rodkevich/tbd/internal/msg"
)

// View ...
func (a Application) View(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	key := r.FormValue("fields")

	parse, err := uuid.Parse(id)
	if err != nil {
		msg.ReturnClientError(w, "bad uuid")
		return
	}

	switch key {
	case "true":
		ticket := ds.TicketWithID(parse, true)
		msg.ReturnJSON(w, ticket)
		return
	}
	ticket := ds.TicketWithID(parse, false)
	msg.ReturnJSON(w, ticket)
}
