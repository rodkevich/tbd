package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rodkevich/tbd/internal/env"
	"github.com/rodkevich/tbd/pkg/app/handlers/tickets"
)

var applicationPort = env.EnvGetOrDefault("APPPORT", "12300")

func Run() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v0/create", tickets.Create).Methods(http.MethodGet)
	router.HandleFunc("/api/v0/search", tickets.Search).Methods(http.MethodGet)
	router.HandleFunc("/api/v0/view", tickets.View).Methods(http.MethodGet)

	log.Println("Starting API server on " + applicationPort)
	if err := http.ListenAndServe(":"+applicationPort, router); err != nil {
		log.Fatal(err)
	}
}
