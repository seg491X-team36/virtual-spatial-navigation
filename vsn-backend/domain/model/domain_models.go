package model

import (
	"time"

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

type Experiment struct {
	Id          uuid.UUID
	Name        string
	Description string
	ArenaId     uuid.UUID
}

type ExperimentResult struct {
	Id           uuid.UUID
	UserId       uuid.UUID // id used to store the experiment data
	ExperimentId uuid.UUID
	Completed    time.Time
}

type ExperimentInput struct {
	ArenaID uuid.UUID
}

type ExperimentUpdateDescriptionInput struct {
	ExperimentID uuid.UUID
	Description  string
}

type ExperimentUpdateNameInput struct {
	ExperimentID uuid.UUID
	Name         string
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
