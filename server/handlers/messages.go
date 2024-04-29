package handlers

import (
	"encoding/json"
	"net/http"

	"secure-messaging-app/server/crypto"
	"secure-messaging-app/server/db"
	"secure-messaging-app/server/models"
)

func SendMessage(w http.ResponseWriter, r *http.Request) {
	var message models.Message
	json.NewDecoder(r.Body).Decode(&message)

	// Encrypt message content
	encryptedContent, _ := crypto.EncryptMessage(message.Content, message.RecipientPublicKey)
	message.Content = encryptedContent

	// Save message to database
	db.SaveMessage(message)

	// Send response
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}
