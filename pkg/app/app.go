package app

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

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
	var wait time.Duration
	flag.DurationVar(
		&wait,
		"graceful-timeout",
		time.Second*15,
		"the duration for graceful stop")
	flag.Parse()

	router := mux.NewRouter()

	router.HandleFunc("/api/v0/create", a.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/v0/list", a.List).Methods(http.MethodGet)
	router.HandleFunc("/api/v0/ticket/{id}", a.View).Methods(http.MethodGet)

	srv := &http.Server{
		Addr:         "0.0.0.0:" + port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}
	log.Println("Starting API server on " + port)
	// start non-blocking
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	// Block until we receive our signal.
	<-c
	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("Stopping API server on " + port)
	os.Exit(0)
}

func (a *Application) SetUp(user, password, dbname string) {
}
