package datasource

import (
	"github.com/google/uuid"

	"github.com/rodkevich/tbd/pkg/tickets"
)

// Datasource ...
type Datasource interface {
	Create(t tickets.Ticket) string
	List(PriceSort, DateSort string) []tickets.Ticket
	TicketWithID(id uuid.UUID, fields bool) (t *tickets.Ticket)
}
