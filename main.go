package main

import (
	"log"

	"github.com/siddhanthpx/phonebook/database"
	"github.com/siddhanthpx/phonebook/routes"
)

func main() {
	err := database.ConnectDB()
	errorHandler("Failed to connect to DB", err)

	routes.SetupRoutes()
}

func errorHandler(msg string, err error) {
	if err != nil {
		log.Fatalf("%s : %s\n", msg, err)
		return
	}
}
