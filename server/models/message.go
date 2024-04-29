package models

type Message struct {
	ID                 string `json:"id"`
	SenderID           string `json:"senderId"`
	RecipientID        string `json:"recipientId"`
	Content            string `json:"content"`
	RecipientPublicKey string `json:"recipientPublicKey"`
}
