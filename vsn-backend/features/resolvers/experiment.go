package resolvers

import (
	"context"
	"errors"

	"github.com/seg491X-team36/vsn-backend/domain/model"
)

type ExperimentResolvers struct {
}

func (r *ExperimentResolvers) Results(ctx context.Context, obj *model.Experiment) ([]model.ExperimentResult, error) {
	return nil, errors.New("not implemented")
}

func (r *ExperimentResolvers) Pending(ctx context.Context, obj *model.Experiment) ([]model.Invite, error) {
	return nil, errors.New("not implemented")
}

func (r *ExperimentResolvers) UsersNotInvited(ctx context.Context, obj *model.Experiment) ([]model.User, error) {
	return nil, errors.New("not implemented")
}
