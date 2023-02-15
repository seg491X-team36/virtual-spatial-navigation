package resolvers

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/seg491X-team36/vsn-backend/domain/model"
)

type QueryResolvers struct {
}

func (q *QueryResolvers) User(ctx context.Context, id uuid.UUID) (*model.User, error) {
	return &model.User{}, errors.New("not implemented")
}

func (q *QueryResolvers) Users(ctx context.Context, state *model.UserAccountState) ([]model.User, error) {
	return []model.User{}, errors.New("not implemented")
}

func (q *QueryResolvers) Invite(ctx context.Context, id uuid.UUID) (*model.Invite, error) {
	return &model.Invite{}, errors.New("not implemented")
}

func (q *QueryResolvers) Invites(ctx context.Context, supervised bool, experiments uuid.UUID) ([]model.Invite, error) {
	return []model.Invite{}, errors.New("not implemented")
}

func (q *QueryResolvers) Experiment(ctx context.Context, id uuid.UUID) (*model.Experiment, error) {
	return &model.Experiment{}, errors.New("not implemented")
}

func (q *QueryResolvers) Experiments(ctx context.Context) ([]model.Experiment, error) {
	return []model.Experiment{}, errors.New("not implemented")
}
