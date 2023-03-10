// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

import (
	"github.com/seg491X-team36/vsn-backend/domain/model"
)

type ExperimentPayload struct {
	Experiment *model.Experiment `json:"experiment"`
	Error      *string           `json:"error"`
}

type InvitePayload struct {
	Invite *model.Invite `json:"invite"`
	Error  *string       `json:"error"`
}

type UserPayload struct {
	User  *model.User `json:"user"`
	Error *string     `json:"error"`
}
