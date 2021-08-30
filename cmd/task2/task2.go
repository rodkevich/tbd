package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/rodkevich/tbd/internal/msg"
	"github.com/rodkevich/tbd/pkg/tickets"
)

const maxTickets = 3

var arr [maxTickets]tickets.Ticket

func main() {
	// add structs:
	for i := 0; i < len(arr); i++ {
		arr[i] = tickets.Ticket{
			// add 1 to skip 0 value:
			OrderNumber: i + 1,
			DateCreated: msg.TimeNowFormatted(),
		}
	}
	// print structs:
	for n, ticket := range arr {
		var out string
		bytes, err := json.Marshal(ticket)
		if err != nil {
			return
		}
		out = fmt.Sprintf(
			msg.StrTemplateTicket,
			n,
			string(bytes))
		log.Println(out)
	}
}
