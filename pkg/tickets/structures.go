package tickets

import (
	"github.com/google/uuid"

	"github.com/rodkevich/tbd/pkg/tickets/types"
)

// Ticket ...
type Ticket struct {
	ID            *uuid.UUID   `json:"id"`
	OrderNumber   uint         `json:"order_number" validate:"gte=1,lte=10"`
	Name          string       `json:"ticket_name" validate:"gte=0,lte=200"`
	PhotoMainLink types.Link   `json:"photo_main_link"`
	PhotoLinks    []types.Link `json:"photo_links" validate:"gte=0,lte=3"`
	Price         struct {
		Currency types.Currency `json:"currency"`
		Current  uint           `json:"current" validate:"gte=0,lte=9223372036854775807"`
		Discount uint           `json:"discount" validate:"gte=0,lte=100"`
		Min      uint           `json:"min" validate:"gte=0,lte=9223372036854775807"`
		Max      uint           `json:"max" validate:"gte=0,lte=9223372036854775807"`
	} `json:"price"`
	Description types.Description `json:"description" validate:"gte=0,lte=1000"`
	PhoneNumber types.Phone       `json:"phone"`
	DateCreated string            `json:"date_created"`
	Deleted     bool              `json:"is_deleted"`
}
