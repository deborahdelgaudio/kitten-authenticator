package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/deborahdelgaudio/kitten_authenticator/handlers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/authorize", handlers.Authorize).Methods("GET")

	http.ListenAndServe(":8080", router)
}
