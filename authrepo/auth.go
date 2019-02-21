package authrepo

import (
	"time"
)

// Session that will hold login user session information.
type Session struct {
	ID          int       `json:"id"`
	Token       string    `json:"token"`
	Email       string    `json:"email"`
	UserID      int       `json:"userId"`
	CreatedDate time.Time `json:"createdDate"`
}
