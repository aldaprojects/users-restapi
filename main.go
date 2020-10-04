package main

import (
	"log"
	"net/http"

	"github.com/aldaprojects/basic-restapi/handlers"
)

func HandleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		handlers.GetUser(w, r)

	case http.MethodPost:
		handlers.PostUser(w, r)

	case http.MethodDelete:
		handlers.DeleteUser(w, r)

	case http.MethodPut:
		handlers.PutUser(w, r)
	}
}

func main() {

	http.HandleFunc("/user", HandleUser)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
