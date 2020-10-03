package db

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"

	"github.com/aldaprojects/basic-restapi/models"
)

type DB struct {
	mux sync.Mutex
}

func read() ([]models.User, error) {
	js, err := ioutil.ReadFile("data/users.json")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var users []models.User
	json.Unmarshal(js, &users)

	return users, nil
}

func write(users []models.User) error {
	json, err := json.Marshal(users)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = ioutil.WriteFile("data/users.json", json, 7777)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (db *DB) Count() (int, error) {
	db.mux.Lock()
	defer db.mux.Unlock()

	users, err := read()
	if err != nil {
		return 0, err
	}
	return len(users), err
}

func (db *DB) Create(user *models.User) error {
	db.mux.Lock()
	defer db.mux.Unlock()

	users, err := read()
	if err != nil {
		return err
	}

	users = append(users, *user)

	return write(users)
}

func (db *DB) Read(id int) (models.User, bool, error) {
	db.mux.Lock()
	defer db.mux.Unlock()

	users, err := read()
	if err != nil {
		return models.User{}, false, err
	}

	for _, user := range users {
		if user.Id == id {
			return user, true, nil
		}
	}

	return models.User{}, false, nil
}

func (db *DB) Update(user models.User) (bool, error) {
	db.mux.Lock()
	defer db.mux.Unlock()

	users, err := read()
	if err != nil {
		return false, err
	}

	for index, userDB := range users {
		if userDB.Id == user.Id {
			users[index] = user
			err := write(users)

			return true, err
		}
	}

	return false, nil

}

func (db *DB) Delete(id int) (bool, error) {
	db.mux.Lock()
	defer db.mux.Unlock()

	users, err := read()
	if err != nil {
		return false, err
	}

	for index, user := range users {
		if user.Id == id {
			users[index] = users[len(users) - 1]
			users = users[:len(users) - 1]

			err := write(users)

			return true, err
		}
	}

	return false, nil
}
