package main

import (
	"fmt"
	"log"

	"github.com/rodkevich/tbd/pkg/tickets"
)

func main() {
	ticket := tickets.Ticket{
		PhoneNumber: "this will fail hardcoded validation",
	}

	// print structs using Stringer:
	fmt.Println(ticket)

	// create Struckturizator
	struckturizator := tickets.NewStruckturizator()

	// validate struct using Struckturizator:
	err := struckturizator.TicketValidation(ticket)
	if err != nil {
		log.Fatal(err) // should fail
	}

}
