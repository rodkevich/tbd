package app

import (
	"net/http"

	"github.com/rodkevich/tbd/internal/msg"
)

// Search ...
func (a Application) Search(w http.ResponseWriter, r *http.Request) {

	msg.ReturnClientError(w, "not yet implemented")

}
