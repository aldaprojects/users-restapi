package db

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"

	"github.com/aldaprojects/basic-restapi/models"
)

type UserDB struct {
	mux sync.Mutex
}

// read reads directly the virtual database and return an array of users
// A successful call returns an array of users an err == nil. Because users.json always has an array.
// First of all users.json has an empty array without users in it
func read() ([]models.User, error) {
	js, err := ioutil.ReadFile("db/users.json")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var users []models.User
	err = json.Unmarshal(js, &users)

	return users, err
}

// write writes a user directly in the virtual database and return an error if there is one.
// A successful call returns err == nil, that means the new user
// was pushed correctly in the array from users.json
func write(users []models.User) error {
	js, err := json.Marshal(users)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = ioutil.WriteFile("db/users.json", js, 7777)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

// Count counts how many users are in the array from users.json
// A successful call returns how many users are in the array and err == nil.
func (db *UserDB) Count() (int, error) {
	db.mux.Lock()
	defer db.mux.Unlock()

	users, err := read()
	if err != nil {
		return 0, err
	}
	return len(users), err
}

// Create creates a user.
// A successful call returns err == nil. First read the database and
// push the new user in the array from arrays.json then write into the database
func (db *UserDB) Create(user *models.User) error {
	db.mux.Lock()
	defer db.mux.Unlock()

	users, err := read()
	if err != nil {
		return err
	}

	users = append(users, *user)

	return write(users)
}

// Read get a user from the virtual database.
// A successful call returns true if user was found and User if User.Id != id.
// If the user was not found it returns false
func (db *UserDB) Read(id int) (models.User, bool, error) {
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

// Update reads the virtual database to get a particular user. Then modify its properties
// and save the user. A successful call return true if the user was modify
// correctly and err == nil.
func (db *UserDB) Update(user models.User) (bool, error) {
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

// Delete reads the virtual database to get a particular user. Then modify the array
// to delete the user and override the array into users.json.
// A successful call returns true if the user was found and deleted correctly and err == nil.
func (db *UserDB) Delete(id int) (bool, error) {
	db.mux.Lock()
	defer db.mux.Unlock()

	users, err := read()
	if err != nil {
		return false, err
	}

	for index, user := range users {
		if user.Id == id {

			// Replace the user with the last user from the array.
			// Then create a slice without the last user.
			users[index] = users[len(users)-1]
			users = users[:len(users)-1]

			err := write(users)

			return true, err
		}
	}

	return false, nil
}
