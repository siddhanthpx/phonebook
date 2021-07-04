package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/siddhanthpx/phonebook/controllers"
)

var PORT = ":8080"

func SetupRoutes() {

	r := mux.NewRouter()
	handlers.LoggingHandler(os.Stdout, r)
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/user", controllers.User).Methods("GET")
	r.HandleFunc("/logout", controllers.Logout).Methods("POST")

	fmt.Printf("Starting server on localhost%s\n", PORT)
	log.Fatal(http.ListenAndServe(PORT, handlers.CORS()(r)))
}
