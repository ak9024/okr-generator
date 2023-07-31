package auth

import "github.com/google/uuid"

type UserModel struct {
	UUID    uuid.UUID `json:"uuid"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	EmailID string    `json:"email_id"`
	Picture string    `json:"picture"`
}
