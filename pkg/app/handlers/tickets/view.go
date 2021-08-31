package tickets

import (
	"errors"
	"net/http"

	"github.com/rodkevich/tbd/internal/msg"
)

// View ...
func View(w http.ResponseWriter, r *http.Request) {
	msg.ReturnServerError(w, errors.New("not yet implemented"))

}
