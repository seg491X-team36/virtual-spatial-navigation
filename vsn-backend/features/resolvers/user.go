package resolvers

import (
	"context"
	"errors"

	"github.com/seg491X-team36/vsn-backend/domain/model"
)

type UserResolvers struct {
}

func (r *UserResolvers) Invites(ctx context.Context, obj *model.User) ([]model.Invite, error) {
	return []model.Invite{}, errors.New("not implemented")
}

func (r *UserResolvers) Results(ctx context.Context, obj *model.User) ([]model.ExperimentResult, error) {
	return []model.ExperimentResult{}, errors.New("not implemented")
}
