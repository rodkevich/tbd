package postgres

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/pgxpool"

	ds "github.com/rodkevich/tbd/pkg/datasource"
	"github.com/rodkevich/tbd/pkg/tickets"
)

// NewDatasourcePG ...
func NewDatasourcePG() (ds.Datasource, error) {
	var config = os.Getenv("POSTGRES_DB")
	pool, poolErr := pgxpool.Connect(context.Background(), config)
	if poolErr != nil {
		log.Fatalf("Unable to connection to database: %v\n", poolErr)
	}
	log.Printf("Connected!")
	return &datasource{pool}, nil
}

// Represents the datasource model
type datasource struct {
	db *pgxpool.Pool
}

func (d datasource) Create(t tickets.Ticket) string {
	panic("implement me")
}

func (d datasource) Search(strings []string) []tickets.Ticket {
	panic("implement me")
}

func (d datasource) ViewByID(id string) tickets.Ticket {
	panic("implement me")
}
