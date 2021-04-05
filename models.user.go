// models.user.go

package main

import (
	"database/sql"
	"errors"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

var DB_user *sql.DB

// For this demo, we're storing the user list in memory
// We also have some users predefined.
// In a real application, this list will most likely be fetched
// from a database. Moreover, in production settings, you should
// store passwords securely by salting and hashing them instead
// of using them as we're doing in this demo
var userList = []user{}

/*var (
	username string
	password string
)*/

// Check if the username and password combination is valid
func isUserValid(username, password string) bool {
	DB_user, err := sql.Open("mysql",
		"root:password@tcp(127.0.0.1:3306)/testDB")
	if err != nil {
		log.Fatal(err)
	}
	defer DB_user.Close()
	rows, err := DB_user.Query("SELECT * FROM Users where Username = '" + username + "' AND Password = '" + password + "'")
	if err != nil {
		// do something here
		log.Fatal(err)
	}
	if !rows.Next() {
		return false
	} else {
		return true
	}
	/*for _, u := range userList {
		if u.Username == username && u.Password == password {
			return true
		}
	}*/
}

// Register a new user with the given username and password
// NOTE: For this demo, we
func registerNewUser(username, password string) (*user, error) {
	DB_user, err := sql.Open("mysql",
		"root:password@tcp(127.0.0.1:3306)/testDB")
	if err != nil {
		log.Fatal(err)
	}
	defer DB_user.Close()
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("The password can't be empty")
	} else if !isUsernameAvailable(username) {
		return nil, errors.New("The username isn't available")
	}
	_, err = DB_user.Exec("INSERT INTO Users(Username,Password) VALUES('" + username + "','" + password + "')")
	if err != nil {
		log.Fatal(err)
	}
	u := user{Username: username, Password: password}

	return &u, nil
}

// Check if the supplied username is available
func isUsernameAvailable(username string) bool {
	/*for _, u := range userList {
		if u.Username == username {
			return false
		}
	}*/
	DB_user, err := sql.Open("mysql",
		"root:password@tcp(127.0.0.1:3306)/testDB")
	if err != nil {
		log.Fatal(err)
	}
	defer DB_user.Close()
	rows, err := DB_user.Query("SELECT * FROM Users where Username = '" + username + "'")
	if err != nil {
		// do something here
		log.Fatal(err)
	}
	if !rows.Next() {
		return true
	} else {
		return false
	}
}
