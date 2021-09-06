package postgres

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
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
	pool, err := pgxpool.Connect(context.Background(), config)

	if err != nil {
		log.Printf("Unable to connect database: %v\n", err)
		return nil, err
	}
	log.Printf("New PG datasource connected to: %v", config)

	return &datasource{pool}, nil
}

func (d datasource) String() string {
	return "Postgres"
}

// Create ...
func (d datasource) Create(t tickets.Ticket) (ticketID string) {
	ctx, cancel := context.WithTimeout(ctxDefault, operationsTimeOut)
	defer cancel()

	stmt = `
		INSERT INTO tickets
		(order_number, ticket_name, photo_main_link, currency, current_price,
		discount, min_price, max_price, description, phone_number,
		is_active, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING id;
		`
	err := d.db.QueryRow(
		ctx, stmt,
		t.OrderNumber, t.Name, t.PhotoMainLink, t.Price.Currency,
		t.Price.Current, t.Price.Discount, t.Price.Min, t.Price.Max,
		t.Description, t.PhoneNumber, t.Active, t.DateCreated).Scan(&ticketID)

	if err != nil {
		log.Println(err)
		return
	}

	batch := &pgx.Batch{}
	for _, photoLink := range t.PhotoLinks {
		stmt2 := `
			INSERT INTO photo_links (ticket_id, link_address)
			VALUES ($1, $2)
			RETURNING link_id;
		`
		batch.Queue(stmt2, ticketID, photoLink)
	}

	br := d.db.SendBatch(ctxDefault, batch)
	ct, err := br.Exec()
	if err != nil {
		log.Println(err)
		return
	}
	if ct.RowsAffected() != 1 {
		log.Printf("%v ticket links: rows affected => %v, want %v", ticketID, ct.RowsAffected(), 1)
	}

	return ticketID
}

// List ...
func (d datasource) List(PriceSort, DateSort string) (items []tickets.Ticket) {
	stmt := `
		WITH some_count AS (
		SELECT id, order_number, ticket_name, photo_main_link, currency, current_price,
		discount, min_price, max_price, description, phone_number,
		is_active, created_at,
		ARRAY_AGG(link_address) photo_links
		FROM tickets
		LEFT JOIN photo_links ON tickets.id = photo_links.ticket_id
		GROUP BY id)
		SELECT *
		FROM some_count
		ORDER BY
		created_at ` + DateSort + `,
		current_price ` + PriceSort + `;
	`
	rows, err := d.db.Query(ctxDefault, stmt)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var t tickets.Ticket
		if err := rows.Scan(
			&t.ID, &t.OrderNumber, &t.Name, &t.PhotoMainLink,
			&t.Price.Currency, &t.Price.Current,
			&t.Price.Discount, &t.Price.Min, &t.Price.Max,
			&t.Description, &t.PhoneNumber, &t.Active, &t.DateCreated, &t.PhotoLinks,
		); err != nil {
			return nil
		}
		items = append(items, t)
	}
	if err := rows.Err(); err != nil {
		log.Printf("err: pg: search: %v\n", err)
		return
	}
	return items
}

// TicketWithID ...
func (d datasource) TicketWithID(id uuid.UUID, fields bool) *tickets.Ticket {
	ctx, cancel := context.WithTimeout(ctxDefault, operationsTimeOut)
	defer cancel()

	stmt := `
		WITH some_count AS (
		SELECT id, order_number, ticket_name, photo_main_link, currency, current_price,
		discount, min_price, max_price, description, phone_number,
		is_active, created_at,
		ARRAY_AGG(link_address) photo_links
		FROM tickets
		LEFT JOIN photo_links ON tickets.id = photo_links.ticket_id
		GROUP BY id)
		SELECT *
		FROM some_count
		WHERE some_count.id = $1;
		`
	var t tickets.Ticket
	row := d.db.QueryRow(ctx, stmt, id)
	if err := row.Scan(
		&t.ID, &t.OrderNumber, &t.Name, &t.PhotoMainLink,
		&t.Price.Currency, &t.Price.Current,
		&t.Price.Discount, &t.Price.Min, &t.Price.Max,
		&t.Description, &t.PhoneNumber, &t.Active, &t.DateCreated, &t.PhotoLinks,
	); err != nil {
		log.Println(err)
		return nil
	}
	// optional part
	if !fields {
		t.PhotoLinks = nil
	}
	return &t
}
