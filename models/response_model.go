package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Response is a format for ResponseWriter. Create just one object.
type Response struct {
	Body  User   `json:"body"`
	Error string `json:"error"`
	Ok    bool   `json:"ok"`
}

// write writes into ResponseWriter interface
func write(resp *Response, w http.ResponseWriter) {
	js, _ := json.Marshal(resp)
	fmt.Fprint(w, string(js))
}

// InternalServerError send an status 500 then write
func (resp *Response) InternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	write(resp, w)
}

// BadRequest send an status 400, add a message then write
func (resp *Response) BadRequest(w http.ResponseWriter, message string) {
	resp.Error = message
	write(resp, w)
}

// NotFound send an status 404, add a message specifying the queryParam then write
func (resp *Response) NotFound(w http.ResponseWriter, queryParam string) {
	resp.Error = "The user with id " + queryParam + " was not found"
	write(resp, w)
}

// StatusOk send an status 200, modify the ok response to true then write
func (resp *Response) StatusOk(w http.ResponseWriter) {
	resp.Ok = true
	write(resp, w)
}
