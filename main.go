package main // import "github.com/eriol/wp24-deities"

import (
	"log"

	"github.com/eriol/wp24-deities/api"
	"github.com/eriol/wp24-deities/database"
)

const dbPath = "deities.sqlite"

func main() {
	err := database.Open(dbPath)
	defer database.Close()
	if err != nil {
		log.Fatal(err)
	}

	api.Serve()
}
