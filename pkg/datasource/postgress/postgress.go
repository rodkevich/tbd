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
type link struct {
	LinkId      int        `json:"link_id,omitempty"`
	TicketId    *uuid.UUID `json:"ticket_id,omitempty"`
	LinkAddress string     `json:"link_address,omitempty"`
}

// NewDatasource ...
func NewDatasource(config string) (ds.Datasource, error) {
	pool, poolErr := pgxpool.Connect(context.Background(), config)

	if poolErr != nil {
		log.Fatalf("Unable to connect database: %v\n", poolErr)
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
			RETURNING id;
			`
	err := d.db.QueryRow(
		ctx, stmt,
		t.OrderNumber, t.Name, t.PhotoMainLink, t.Price.Currency,
		t.Price.Current, t.Price.Discount, t.Price.Min, t.Price.Max,
		t.Description, t.PhoneNumber, t.Active, t.DateCreated).Scan(&ticketID)

	if err != nil {
		log.Println(err)
	}

	for _, photoLink := range t.PhotoLinks {
		stmt2 := `
			INSERT INTO photo_links (ticket_id, link_address)
			VALUES ($1, $2);
			`
		err = d.db.QueryRow(ctx, stmt2, ticketID, photoLink).Scan()
		if err != nil {
			log.Println(err)
		}
	}
	return
}

// List ...
func (d datasource) List(PriceSort, DateSort string) (items []tickets.Ticket) {
	stmt := `
	SELECT
	*
	FROM
	tickets
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
			&t.ID,
			&t.OrderNumber,
			&t.Name,
			&t.PhotoMainLink,
			&t.Price.Currency,
			&t.Price.Current,
			&t.Price.Discount,
			&t.Price.Min,
			&t.Price.Max,
			&t.Description,
			&t.PhoneNumber,
			&t.Active,
			&t.DateCreated,
		); err != nil {
			return nil
		}

		getTicketLinks := `
		SELECT  link_address
		FROM photo_links
		WHERE ticket_id = $1;
		`
		rows, err := d.db.Query(ctxDefault, getTicketLinks, t.ID)
		if err != nil {
			log.Println(err)
		}
		var link string
		for rows.Next() {
			if err := rows.Scan(
				&link,
			); err != nil {
				log.Println(err)
			}
			t.PhotoLinks = append(t.PhotoLinks, link)
		}
		items = append(items, t)
	}
	log.Printf("items: %v\n", items)

	if err := rows.Err(); err != nil {
		log.Printf("err: pg: search: %v\n", err)
		return
	}
	return items
}

// TicketWithID ...
func (d datasource) TicketWithID(id string) *tickets.Ticket {
	print(id)
	ctx, cancel := context.WithTimeout(ctxDefault, operationsTimeOut)
	defer cancel()
	stmt := `
			WITH some_count AS (
			SELECT
			id,
			ARRAY_AGG (link_address) photo_links
			FROM
			tickets
			LEFT JOIN photo_links ON tickets.id=photo_links.ticket_id
			GROUP BY
			id)
			SELECT * FROM some_count
			WHERE id=$1;
			`
	var t tickets.Ticket
	err := d.db.QueryRow(ctx, stmt, id).Scan(
		&t.ID, &t.PhotoLinks,
		// &t.OrderNumber, &t.Name, &t.PhotoMainLink, &t.Price.Currency,
		// &t.Price.Current, &t.Price.Discount, &t.Price.Min, &t.Price.Max,
		// &t.Description, &t.PhoneNumber, &t.Active, &t.DateCreated
	)

	if err != nil {
		log.Printf("err: pg: view by id: stmt: %v\n", err)
		return nil
	}

	return &t
}

// // TicketWithID ...
// func (d datasource) TicketWithID(id string) *tickets.Ticket {
// 	ctx, cancel := context.WithTimeout(ctxDefault, operationsTimeOut)
// 	defer cancel()
// 	stmt := `
// 		SELECT
// 		(id, order_number, ticket_name, photo_main_link, currency, current_price,
// 		discount, min_price, max_price, description, phone_number,
// 		is_active, created_at)
// 		from tickets left join photo_links on tickets.id=photo_links.ticket_id
// 		WHERE id=$1
// 		AND is_active=true
// 		`
// 	var t tickets.Ticket
// 	err := d.db.QueryRow(ctx, stmt, id).Scan(
// 		&t.ID,
// 		&t.OrderNumber, &t.Name, &t.PhotoMainLink, &t.Price.Currency,
// 		&t.Price.Current, &t.Price.Discount, &t.Price.Min, &t.Price.Max,
// 		&t.Description, &t.PhoneNumber, &t.Active, &t.DateCreated)
//
// 	if err != nil {
// 		log.Printf("err: pg: view by id: stmt: %v\n", err)
// 		return nil
// 	}
// 	return &t
// }
//
