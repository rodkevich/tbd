package tickets

import (
	"time"

	"github.com/google/uuid"

	"github.com/rodkevich/tbd/pkg/tickets/types"
)

// Ticket ...
type Ticket struct {
	ID            *uuid.UUID `json:"id"`
	OrderNumber   uint       `json:"order_number" validate:"gte=1,lte=10"`
	Name          string     `json:"ticket_name" validate:"gte=0,lte=200"`
	PhotoMainLink types.Link `json:"photo_main_link"`
	PhotoLinks    []*string  `json:"photo_links,omitempty" validate:"lte=3"`
	Price         struct {
		Currency types.Currency `json:"currency"`
		Current  float64        `json:"current_price" validate:"gte=0,lte=9223372036854775807"`
		Discount uint           `json:"discount" validate:"gte=0,lte=100"`
		Min      float64        `json:"min_price" validate:"gte=0,lte=9223372036854775807"`
		Max      float64        `json:"max_price" validate:"gte=0,lte=9223372036854775807"`
	} `json:"price"`
	Description types.Description `json:"description,omitempty" validate:"gte=0,lte=1000"`
	PhoneNumber types.Phone       `json:"phone_number"`
	DateCreated time.Time         `json:"created_at"`
	Active      bool              `json:"is_active"`
}
