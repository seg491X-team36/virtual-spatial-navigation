package model

import (
	"github.com/google/uuid"
)

type UserAccountState string

const (
	REGISTERED = UserAccountState("REGISTERED")
	REJECTED   = UserAccountState("REJECTED")
	ACCEPTED   = UserAccountState("ACCEPTED")
)

type User struct {
	Id     uuid.UUID
	Email  string
	State  UserAccountState // registered, rejected, accepted
	Source string           // google form, researcher
}

type Invite struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	ExperimentID uuid.UUID
	Supervised   bool
}

type Arena struct {
	Id uuid.UUID
}

type InviteInput struct {
	UserID       uuid.UUID
	ExperimentID uuid.UUID
	Supervised   bool
}

type UserSelectionInput struct {
	UserID uuid.UUID
	Accept bool
}
