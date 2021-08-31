package tickets

import (
	"net/http"

	"github.com/rodkevich/tbd/internal/msg"
	"github.com/rodkevich/tbd/pkg/tickets"
)

// Create ...
func Create(w http.ResponseWriter, r *http.Request) {
	time := msg.TimeNowFormatted()
	t := tickets.Ticket{DateCreated: time}
	msg.ReturnJSON(w, t)
}
