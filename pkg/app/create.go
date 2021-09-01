package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/rodkevich/course-rest/lib"
	"github.com/rodkevich/tbd/internal/msg"
	"github.com/rodkevich/tbd/pkg/tickets"
)

// Create ...
func (a Application) Create(w http.ResponseWriter, r *http.Request) {
	var t tickets.Ticket
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg.ReturnServerError(w, err)
		return
	}
	if err := json.Unmarshal(b, &t); err != nil {
		lib.ReturnInternalError(w, err)
		return
	}
	log.Println(t.String())

	err = a.TicketValidation(t)
	if err != nil {
		msg.ReturnClientError(w, "Ticket-Validation-Error")
		return
	}
	uuid := ds.Create(t)
	msg.ReturnJSON(w, uuid)
}
