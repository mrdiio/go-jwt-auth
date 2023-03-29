package response

import "github.com/google/uuid"

type UserResponse struct {
	ID       uuid.UUID `json:"id,omitempty"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Email    *string   `json:"email"`
}
