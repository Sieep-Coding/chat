package db

import (
	"database/sql"

	_ "github.com/lib/pq"

	"secure-messaging-app/server/config"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("postgres", config.DatabaseURL)
	if err != nil {
		panic(err)
	}
}

func CreateUser(user models.User) {
	// Insert user into the database
	// ...
}

func GetUserByUsername(username string) models.User {
	// Retrieve user from the database by username
	// ...
}

func SaveMessage(message models.Message) {
	// Save message to the database
	// ...
}
