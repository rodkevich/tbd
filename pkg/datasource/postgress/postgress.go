package postgres

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	ds "github.com/rodkevich/tbd/pkg/datasource"
	"github.com/rodkevich/tbd/pkg/tickets"
)

var (
	stmt              string
	rows              pgx.Rows
	row               pgx.Row
	ctxDefault        = context.Background()
	operationsTimeOut = 3 * time.Second
)

// Represents the datasource model
type datasource struct {
	db *pgxpool.Pool
}

// NewDatasource ...
func NewDatasource(config string) (ds.Datasource, error) {
	pool, poolErr := pgxpool.Connect(context.Background(), config)

	if poolErr != nil {
		log.Fatalf("Unable to connection to database: %v\n", poolErr)
	}
	log.Printf("Connected!")

	return &datasource{pool}, nil
}

func (d datasource) String() string {
	return "Postgres"
}

// Create ...
func (d datasource) Create(t tickets.Ticket) (ticketID string) {
	ctx, cancel := context.WithTimeout(ctxDefault, operationsTimeOut)
	defer cancel()

	stmt = `INSERT INTO tickets
			(order_number, ticket_name, photo_main_link, currency, current_price,
			discount, min_price, max_price, description, phone_number,
			is_active, created_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
			RETURNING ticket_id;
			`
	err := d.db.QueryRow(
		ctx,
		stmt,
		t.OrderNumber, t.Name, t.PhotoMainLink, t.Price.Currency,
		t.Price.Current, t.Price.Discount, t.Price.Min, t.Price.Max,
		t.Description, t.PhoneNumber, t.Active, t.DateCreated).Scan(&ticketID)

	if err != nil {
		log.Println(err)
	}

	return
}

// Search ...
func (d datasource) Search(strings []string) []tickets.Ticket {

	return nil
}

// ViewByID ...
func (d datasource) ViewByID(id string) (t *tickets.Ticket) {
	ctx, cancel := context.WithTimeout(ctxDefault, operationsTimeOut)
	defer cancel()
	stmt := `
		SELECT *
		FROM tickets
		WHERE ticket_id=$1
		AND is_active=true
		`
	var tic = tickets.Ticket{}
	// err := d.db.QueryRow(ctx, stmt, id).Scan(
	// 	// &tic.ID,
	// 	// &tic.OrderNumber,
	// 	// &tic.Name,
	// 	// &tic.Description,
	// 	// &tic.Active,
	// 	// &tic.DateCreated,
	// 	// &tic.Price.Current,
	// &tic)
	err := d.db.QueryRow(ctx, stmt, id).Scan(&t)

	log.Println(tic)
	if err != nil {
		log.Printf("err: pg: view by id: stmt: %v\n", err)
		return nil
	}

	return
}
