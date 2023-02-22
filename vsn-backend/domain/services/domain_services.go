package services

import (
	"context"

	"github.com/seg491X-team36/vsn-backend/domain/model"
)

type UserService interface {
	Register(ctx context.Context, email string, source string) (model.User, error)
	Select(ctx context.Context, input []model.UserSelectInput) []model.User // list of users that were updated
}

type InviteService interface {
	Send(ctx context.Context, input model.InviteInput) (model.Invite, error)
}
