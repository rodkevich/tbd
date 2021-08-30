package tickets

import (
	"encoding/json"
	"fmt"

	"github.com/rodkevich/tbd/internal/msg"
)

func (t Ticket) String() (s string) {
	bytes, err := json.Marshal(t)
	if err != nil {
		return
	}
	s = fmt.Sprintf(
		msg.StrTemplateTicket,
		t.OrderNumber,
		string(bytes))
	return
}
