package main

import (
	"log"
	"net/http"

	"github.com/aldaprojects/basic-restapi/queries"
)

func HandleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		queries.GetUser(w, r)

	case http.MethodPost:
		queries.PostUser(w, r)

	case http.MethodDelete:
		queries.DeleteUser(w, r)

	case http.MethodPut:
		queries.PutUser(w, r)
	}
}

func main() {
	http.HandleFunc("/user", HandleUser)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
