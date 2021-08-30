package main

import (
	"log"

	"github.com/rodkevich/tbd/pkg/tickets"
	"github.com/rodkevich/tbd/pkg/tickets/types"
)

func main() {
	ticket := tickets.Ticket{
		OrderNumber: 1,
		PhoneNumber: `(123) 456-7890`,
	}
	// задание: одна из структур должна быть встроена в другую
	ticket.Price.Currency = types.BYN
	ticket.Price.Current = 100500

	// задание: добавить метод для работы с основной структурой
	method := ticket.String
	log.Println(method())

	// задание: добавить функцию для работы с основной структурой
	function := tickets.TicketValidation
	err := function(ticket)
	if err != nil {
		log.Fatal(err)
	}
}
