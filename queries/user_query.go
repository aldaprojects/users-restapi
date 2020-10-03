package queries

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/aldaprojects/basic-restapi/db"
	"github.com/aldaprojects/basic-restapi/models"
	"github.com/aldaprojects/basic-restapi/response"
)

func GetUser(w http.ResponseWriter, r *http.Request)  {

	var (
		resp = response.Response{}
		userDB = db.DB{}
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

func PostUser(w http.ResponseWriter, r *http.Request)  {
	var (
		resp = response.Response{}
		userDB = db.DB{}
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

func DeleteUser(w http.ResponseWriter, r *http.Request)  {
	var (
		resp = response.Response{}
		userDB = db.DB{}
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

func PutUser(w http.ResponseWriter, r *http.Request)  {
	var (
		resp = response.Response{}
		userDB = db.DB{}
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

