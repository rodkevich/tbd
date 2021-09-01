package tickets

// NOTE: could use interface and make the struct private
// 	it will be: type struckturizator struct{}
//	but changed it for demo of embedding in task5 application
//
// type Struckturizator interface {
// 	TicketValidation(t Ticket, opts ...ValidationOption) (err error)
// }

// Struckturizator ..
type Struckturizator struct {
}

// NewStruckturizator ...
func NewStruckturizator() *Struckturizator {
	return &Struckturizator{}
}

// TicketValidation small trick :)
func (s *Struckturizator) TicketValidation(t Ticket, opts ...ValidationOption) (err error) {
	return TicketValidation(t, opts...)
}
