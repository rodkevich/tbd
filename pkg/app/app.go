package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rodkevich/tbd/pkg/app/handlers/tickets"
)

func Run() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v0/create", tickets.Create).Methods(http.MethodGet)
	router.HandleFunc("/api/v0/search", tickets.Search).Methods(http.MethodGet)
	router.HandleFunc("/api/v0/view", tickets.View).Methods(http.MethodGet)

	log.Println("Starting API server on 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
