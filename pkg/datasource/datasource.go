package datasource

import "github.com/rodkevich/tbd/pkg/tickets"

// Datasource ...
type Datasource interface {
	Create(t tickets.Ticket) string
	List(PriceSort, DateSort string) []tickets.Ticket
	TicketWithID(id string) (t *tickets.Ticket)
}
