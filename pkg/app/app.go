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

// Application ...
type Application struct {
	Router                   *mux.Router
	*tickets.Struckturizator // Embedding task3
}

// Run ...
func (a *Application) Run(port string, config string) {

	// Init datasource
	ds, err = postgres.NewDatasource(config)
	if err != nil {

		time.Sleep(time.Second * 3)
		log.Println("trying to reconnect to data-source in 3 sec")
		ds, err = postgres.NewDatasource(config)

		if err != nil {
			defer func() {
				if err := recover(); err != nil {
					log.Println("err: data-source isn't ready")
				}
			}()
			log.Println(err)
			panic(err)
		}
	}

	// Run sever instance in goroutine
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
	// Start non-blocking
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	// Block until we receive our signal
	<-c
	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("Stopping API server on " + port)
	os.Exit(0)
}
