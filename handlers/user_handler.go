package handlers

import (
	"encoding/json"
	"github.com/aldaprojects/basic-restapi/db"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/aldaprojects/basic-restapi/models"
)

// GetUser gets a particular user from the database.
// Request must have an 'id' query param.
// A successful call sends to client a json with the format response and the user
func GetUser(w http.ResponseWriter, r *http.Request) {

	var (
		resp   = models.Response{}
		userDB = db.UserDB{}
	)

	ids, ok := r.URL.Query()["id"]
	if !ok {
		resp.BadRequest(w, "query param 'id' is required")
	}

	id, _ := strconv.Atoi(ids[0])
	user, ok, err := userDB.Read(id)
	if err != nil {
		resp.InternalServerError(w)
		return
	}

	if !ok {
		resp.NotFound(w, ids[0])
		return
	}

	resp.Body = user
	resp.StatusOk(w)
}

// PostUser creates a new user. Request must have a json into the body.
// The json must contain key username and key password.
// A successful call post a user into the database
func PostUser(w http.ResponseWriter, r *http.Request) {
	var (
		resp   = models.Response{}
		userDB = db.UserDB{}
	)

	body, _ := ioutil.ReadAll(r.Body)

	var user models.User
	json.Unmarshal(body, &user)

	if user.Username == "" {
		resp.BadRequest(w, "key 'username' is required")
		return
	}

	if user.Password == "" {
		resp.BadRequest(w, "key 'password' is required")
		return
	}

	count, err := userDB.Count()
	if err != nil {
		resp.InternalServerError(w)
		return
	}

	user.Id = count

	err = userDB.Create(&user)
	if err != nil {
		resp.InternalServerError(w)
		return
	}

	resp.Body = user
	resp.StatusOk(w)
}

// DeleteUser deletes a particular user from the database
// Request must have an 'id' query param
// A successful call remove a user from the database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var (
		resp   = models.Response{}
		userDB = db.UserDB{}
	)

	ids, ok := r.URL.Query()["id"]
	if !ok {
		resp.BadRequest(w, "query param 'id' is required")
		return
	}

	id, _ := strconv.Atoi(ids[0])
	ok, err := userDB.Delete(id)
	if err != nil {
		resp.InternalServerError(w)
		return
	}

	if !ok {
		resp.NotFound(w, ids[0])
		return
	}
	resp.StatusOk(w)
}

// PutUser modify a user that already exists.
// Request must have an 'id' query param and a json into the body
// with key password and key username. A successful call modify a user from the database
func PutUser(w http.ResponseWriter, r *http.Request) {
	var (
		resp   = models.Response{}
		userDB = db.UserDB{}
	)

	ids, ok := r.URL.Query()["id"]
	if !ok {
		resp.BadRequest(w, "query param 'id' is required")
		return
	}

	body, _ := ioutil.ReadAll(r.Body)

	var user models.User
	json.Unmarshal(body, &user)

	if user.Username == "" {
		resp.BadRequest(w, "key username not found")
		return
	}

	if user.Password == "" {
		resp.BadRequest(w, "key password not found")
		return
	}

	id, _ := strconv.Atoi(ids[0])
	user.Id = id

	ok, err := userDB.Update(user)
	if err != nil {
		resp.InternalServerError(w)
		return
	}
	if !ok {
		resp.NotFound(w, ids[0])
		return
	}

	resp.StatusOk(w)
}
