package tickets

// Struckturizator ...
type Struckturizator interface {
	TicketValidation(t Ticket, opts ...ValidationOption) (err error)
}

// Private structure
type struckturizator struct {
}

// NewStruckturizator ...
func NewStruckturizator() *struckturizator {
	return &struckturizator{}
}

// TicketValidation small trick :)
func (s *struckturizator) TicketValidation(t Ticket, opts ...ValidationOption) (err error) {
	return TicketValidation(t, opts...)
}
