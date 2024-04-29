package main

import (
	"log"
	"net/http"

	"secure-messaging-app/server/config"
	"secure-messaging-app/server/db"
	"secure-messaging-app/server/handlers"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Initialize database connection
	db.InitDB()

	// Register routes
	http.HandleFunc("/api/auth/register", handlers.RegisterUser)
	http.HandleFunc("/api/auth/login", handlers.LoginUser)
	http.HandleFunc("/api/messages", handlers.SendMessage)

	// Start server
	log.Println("Server started on port", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
