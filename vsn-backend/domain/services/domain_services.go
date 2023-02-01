package services

import (
	"github.com/seg491X-team36/vsn-backend/domain/model"
)

type UserService interface {
	Register(emails string) error
	Select(input model.UserSelectionInput) error
}

type InviteService interface {
	Create(input model.InviteInput) (model.Invite, error)
}

type ExperimentService interface {
	Create(input model.ExperimentInput) (model.Experiment, error)
	UpdateName(input model.ExperimentUpdateNameInput) (model.Experiment, error)
	UpdateDescription(input model.ExperimentUpdateDescriptionInput) (model.Experiment, error)
}
