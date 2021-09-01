package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rodkevich/tbd/pkg/datasource"
	postgres "github.com/rodkevich/tbd/pkg/datasource/postgress"
	"github.com/rodkevich/tbd/pkg/tickets"
)

var (
	ds  datasource.Datasource
	err error
)

func init() {
	var config = "postgresql://postgres:postgres@localhost:5432/postgres"
	ds, err = postgres.NewDatasource(config)
	if err != nil {
		panic(err)
	}
}

// Application ...
type Application struct {
	Router                   *mux.Router
	*tickets.Struckturizator // embedding task3
}

// Run ...
func (a *Application) Run(port string) {
	router := mux.NewRouter()

	router.HandleFunc("/api/v0/create", a.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/v0/search", a.Search).Methods(http.MethodPost)
	router.HandleFunc("/api/v0/ticket/{id}", a.View).Methods(http.MethodGet)
	router.HandleFunc("/api/v0/ticket/{id}", a.View).Methods(http.MethodGet).Queries(
		"fields", "{fields}").HandlerFunc(a.View).Name("View")

	log.Println("Starting API server on " + port)
	if err = http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}

func (a *Application) SetUp(user, password, dbname string) {
}
