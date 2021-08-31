package tickets

import (
	"net/http"

	"github.com/rodkevich/tbd/internal/msg"
	"github.com/rodkevich/tbd/pkg/tickets"
)

func Create(w http.ResponseWriter, r *http.Request) {
	_ = tickets.Ticket{}
	msg.ReturnClientError(w, "not yet implemented")
}
