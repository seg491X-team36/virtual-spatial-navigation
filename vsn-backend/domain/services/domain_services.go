package services

import (
	"github.com/seg491X-team36/vsn-backend/domain/model"
)

type UserService interface {
	Register(email string, source string) (model.User, error)
	Select(input []model.UserSelectInput) []model.User // list of users that were updated
}

type InviteService interface {
	Send(input model.InviteInput) (model.Invite, error)
}
