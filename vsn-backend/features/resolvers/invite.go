package resolvers

import (
	"context"
	"errors"

	"github.com/seg491X-team36/vsn-backend/domain/model"
)

type InviteResolvers struct {
}

func (r *InviteResolvers) User(ctx context.Context, obj *model.Invite) (model.User, error) {
	return model.User{}, errors.New("not implemented")
}

func (r *InviteResolvers) Experiment(ctx context.Context, obj *model.Invite) (model.Experiment, error) {
	return model.Experiment{}, errors.New("not implemented")
}
