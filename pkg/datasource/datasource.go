package datasource

import "github.com/rodkevich/tbd/pkg/tickets"

// Datasource ...
type Datasource interface {
	Create(t tickets.Ticket) string
	Search([]string) []tickets.Ticket
	ViewByID(id string) (t *tickets.Ticket)
}
