package main

import (
	"github.com/aldaprojects/basic-restapi/services"
	"log"
	"net/http"
)

func HandleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		services.GetUser(w, r)

	case http.MethodPost:
		services.PostUser(w, r)

	case http.MethodDelete:
		services.DeleteUser(w, r)

	case http.MethodPut:
		services.PutUser(w, r)
	}
}

func main() {
	http.HandleFunc("/user", HandleUser)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
