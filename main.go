package main

import (
	"fmt"
	"log"

	"github.com/siddhanthpx/phonebook/database"
)

func main() {
	err := database.ConnectDB()
	errorHandler("Failed to connect to DB", err)

	fmt.Printf("Successfully connected to database: %s\n", database.DB.Name())
}

func errorHandler(msg string, err error) {
	if err != nil {
		log.Fatalf("%s : %s\n", msg, err)
		return
	}
}
