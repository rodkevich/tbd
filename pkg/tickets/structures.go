package tickets

import (
	"github.com/google/uuid"

	"github.com/rodkevich/tbd/pkg/tickets/types"
)

// Ticket ...
type Ticket struct {
	ID            *uuid.UUID   `json:"id"`
	UserID        *uuid.UUID   `json:"user_id"`
	OrderNumber   int          `json:"order_number" validate:"gte=1,lte=10"`
	Name          string       `json:"name" validate:"gte=0,lte=200"`
	PhotoMainLink types.Link   `json:"photo_main_link"`
	PhotoLinks    []types.Link `json:"photo_links" validate:"gte=0,lte=3"`
	Price         struct {
		Currency types.Currency `json:"currency"`
		Current  int            `json:"current" validate:"gte=0,lte=9223372036854775807"`
		Discount int            `json:"discount" validate:"gte=0,lte=100"`
		Min      int            `json:"min" validate:"gte=0,lte=9223372036854775807"`
		Max      int            `json:"max" validate:"gte=0,lte=9223372036854775807"`
	} `json:"price"`
	Description types.Description `json:"description" validate:"gte=0,lte=1000"`
	PhoneNumber types.Phone       `json:"phone"`
	DateCreated string            `json:"date_created"`
}
