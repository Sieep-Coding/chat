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
	err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Initialize database connection
	db.InitDB()

	// Register routes
	http.HandleFunc("/api/auth/register", handlers.RegisterUser)
	http.HandleFunc("/api/auth/login", handlers.LoginUser)
	http.HandleFunc("/api/messages", handlers.SendMessage)

	// Start server
	log.Printf("Server started on port %s", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
