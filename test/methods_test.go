package test

import (
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/rodkevich/tbd/pkg/tickets"
	"github.com/rodkevich/tbd/pkg/tickets/types"
)

func TestTicket_String(t1 *testing.T) {
	type fields struct {
		ID            *uuid.UUID
		OrderNumber   uint
		Name          string
		PhotoMainLink types.Link
		PhotoLinks    []string
		Price         struct {
			Currency types.Currency `json:"currency"`
			Current  float64        `json:"current_price" validate:"gte=0,lte=9223372036854775807"`
			Discount uint           `json:"discount" validate:"gte=0,lte=100"`
			Min      float64        `json:"min_price" validate:"gte=0,lte=9223372036854775807"`
			Max      float64        `json:"max_price" validate:"gte=0,lte=9223372036854775807"`
		}
		Description types.Description
		PhoneNumber types.Phone
		DateCreated time.Time
		Active      bool
	}
	tests := []struct {
		name   string
		fields fields
		wantS  string
	}{
		{
			"1",
			fields{
				OrderNumber:   3,
				PhoneNumber:   `(123) 456-7890`,
				PhotoMainLink: "http://www.example.com",
			},
			`3 | {"id":null,"order_number":3,"ticket_name":"","photo_main_link":"http://www.example.com","price":{"currency":0,"current_price":0,"discount":0,"min_price":0,"max_price":0},"phone_number":"(123) 456-7890","created_at":"0001-01-01T00:00:00Z","is_active":false}`,
		},
		{
			"2",
			fields{},
			`0 | {"id":null,"order_number":0,"ticket_name":"","photo_main_link":"","price":{"currency":0,"current_price":0,"discount":0,"min_price":0,"max_price":0},"phone_number":"","created_at":"0001-01-01T00:00:00Z","is_active":false}`,
		},
		{
			"3",
			fields{
				OrderNumber:   1,
				Name:          "Name Example",
				PhotoMainLink: "http://www.example.com",
				PhotoLinks: []string{
					"http://www.example.com/a#",
					"https://www.example.com/b?a=b%20c",
					"ws://www.example.com/websocket",
				},
				Description: "Description Here",
				PhoneNumber: "+91 (123) 456-7890",
				Active:      true,
			},
			`1 | {"id":null,"order_number":1,"ticket_name":"Name Example","photo_main_link":"http://www.example.com","photo_links":["http://www.example.com/a#","https://www.example.com/b?a=b%20c","ws://www.example.com/websocket"],"price":{"currency":0,"current_price":0,"discount":0,"min_price":0,"max_price":0},"description":"Description Here","phone_number":"+91 (123) 456-7890","created_at":"0001-01-01T00:00:00Z","is_active":true}`,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := tickets.Ticket{
				ID:            tt.fields.ID,
				OrderNumber:   tt.fields.OrderNumber,
				Name:          tt.fields.Name,
				PhotoMainLink: tt.fields.PhotoMainLink,
				PhotoLinks:    tt.fields.PhotoLinks,
				Price:         tt.fields.Price,
				Description:   tt.fields.Description,
				PhoneNumber:   tt.fields.PhoneNumber,
				DateCreated:   tt.fields.DateCreated,
				Active:        tt.fields.Active,
			}
			if gotS := t.String(); gotS != tt.wantS {
				t1.Errorf("String() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}
