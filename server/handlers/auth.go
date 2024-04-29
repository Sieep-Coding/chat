package handlers

import (
	"encoding/json"
	"net/http"

	"secure-messaging-app/server/crypto"
	"secure-messaging-app/server/db"
	"secure-messaging-app/server/models"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	// Hash the password
	hashedPassword, _ := crypto.HashPassword(user.Password)
	user.Password = hashedPassword

	// Save user to database
	db.CreateUser(user)

	// Generate session token
	token := crypto.GenerateToken(user.ID)

	// Send response
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	// Retrieve user from database
	dbUser := db.GetUserByUsername(user.Username)

	// Verify password
	if !crypto.VerifyPassword(user.Password, dbUser.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate session token
	token := crypto.GenerateToken(dbUser.ID)

	// Send response
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
