package model

import (
	"time"

	"github.com/google/uuid"
)

type UserAccountState string

const (
	UserAccountStateRegistered UserAccountState = "REGISTERED"
	UserAccountStateRejected   UserAccountState = "REJECTED"
	UserAccountStateAccepted   UserAccountState = "ACCEPTED"
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
	ArenaId     int
}

type ExperimentResult struct {
	Id           uuid.UUID
	UserId       uuid.UUID // id used to store the experiment data
	ExperimentId uuid.UUID
	Completed    time.Time
}

type ExperimentInput struct {
	ArenaID string
}

type ExperimentUpdateDescriptionInput struct {
	ExperimentID string
	Description  string
}

type ExperimentUpdateNameInput struct {
	ExperimentID string
	Name         string
}

type InviteInput struct {
	UserID       string
	ExperimentID string
	Supervised   bool
}

type UserSelectionInput struct {
	UserID string
	Accept bool
}
