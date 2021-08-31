package main

import (
	"log"

	"github.com/rodkevich/tbd/pkg/tickets"
	"github.com/rodkevich/tbd/pkg/tickets/types"
)

// StructureValidationSolution ...
func main() {
	ticket := tickets.Ticket{OrderNumber: 1}
	ticket.Price.Current = 100
	ticket.Description = `Description`
	ticket.Name = `Name`
	ticket.PhotoLinks = []types.Link{
		"http://www.example.com/a#",
		"https://www.example.com/b?a=b%20c",
		"ws://www.example.com/websocket",
	}
	ticket.PhoneNumber = `+91 (123) 456-7890`

	validate := tickets.TicketValidation

	err := validate(
		ticket,
		// disable opts:
		tickets.WithoutStructTags(),
		// enable opts:
		tickets.WithNameCheck(),
		tickets.WithDescriptionCheck(),
		tickets.WithPhotoLinksCheck(),
	)
	if err != nil {
		log.Fatal(err)
	}
}
