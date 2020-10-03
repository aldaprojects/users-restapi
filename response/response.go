package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aldaprojects/basic-restapi/models"
)

type Response struct {
	Body  models.User `json:"body"`
	Error string      `json:"error"`
	Ok    bool        `json:"ok"`
}

func write(resp *Response, w http.ResponseWriter)  {
	js, _ := json.Marshal(resp)
	fmt.Fprint(w, string(js))
}

func (resp *Response) InternalServerError(w http.ResponseWriter)  {
	w.WriteHeader(http.StatusInternalServerError)
	write(resp, w)
}

func (resp *Response) BadRequest(w http.ResponseWriter, message string) {
	resp.Error = message
	write(resp, w)
}

func (resp *Response) NotFound(w http.ResponseWriter, queryParam string)  {
	resp.Error = "The user with id " + queryParam + " was not found"
	write(resp, w)
}

func (resp *Response) StatusOk(w http.ResponseWriter) {
	resp.Ok = true
	write(resp, w)
}