package main

import (
	"os"

	"github.com/rodkevich/tbd/pkg/app"
)

func main() {
	a := app.Application{}
	a.Run(os.Getenv("APP_API_PORT"))
}
