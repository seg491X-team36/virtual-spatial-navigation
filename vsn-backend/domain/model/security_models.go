package model

import "github.com/google/uuid"

type UserClaims struct {
	UserId uuid.UUID `json:"userId"`
}

type AdminClaims struct {
	Email string `json:"email"`
}
