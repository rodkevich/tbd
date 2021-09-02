package main

import (
	"github.com/rodkevich/tbd/internal/env"
	"github.com/rodkevich/tbd/pkg/app"
)

func main() {

	appPort := env.GetOrDefault(
		// make it run locally & in docker
		"APP_API_PORT",
		"12300",
	)
	dataSourceURL := env.GetOrDefault(
		// for docker db on localhost
		"DATASOURCE_URL",
		"postgresql://postgres:postgres@localhost:5432/postgres",
	)
	a := app.Application{}
	a.Run(appPort, dataSourceURL)
}

/*
NOTE: to get uuid like - "c0c31f94-d14d-4c5b-81ef-1058d5906f70" use :

curl -X POST 'localhost:12300/api/v0/create' \
--header 'Content-Type: application/json' \
--data-raw '{
    "order_number": 1,
    "ticket_name": "Name Example",
    "photo_main_link": "http://www.example.com",
    "photo_links": [
        "http://www.example.com/a#",
        "https://www.example.com/b?a=b%20c",
        "ws://www.example.com/websocket"
    ],
    "price": {
        "currency": 251,
        "current_price": 150.1,
        "discount": 10,
        "min_price": 22.22,
        "max_price": 33.33
    },
    "description": "Description Here",
    "phone_number": "+91 (123) 456-7890",
    "created_at": "0021-01-01T00:00:00Z",
    "is_active": true
}'

NOTE: to get ticket paste new uuid & call :
curl -X GET 'localhost:12300/api/v0/ticket/c0c31f94-d14d-4c5b-81ef-1058d5906f70?fields=true'

NOTE: to get list of tickets call :
curl -X GET 'localhost:12300/api/v0/list?price=DESC&date=ASC'

*/
