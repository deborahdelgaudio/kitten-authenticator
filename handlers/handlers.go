package handlers

import (
	"fmt"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	username := r.FormValue("username")
	msg := fmt.Sprintf("Kitten says: Hello %s! Can I call you %s?", name, username)
	fmt.Fprintf(w, "%s", msg)
	// Save user's data
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	msg := fmt.Sprintf("Kitten says: your username is %s, and your password is %s", username, password)
	fmt.Fprintf(w, "%s", msg)
	// To Do:
	// Check username and password are valid
	// Create JWT token
	// Sign JWT token with privare key
	// Return JWT within response header

}

func Authorize(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	msg := fmt.Sprintf("Kitten says: Do you want to access Kitten's house with %s", token)
	fmt.Fprintf(w, "%s", msg)
	// Validate JWT token
}

