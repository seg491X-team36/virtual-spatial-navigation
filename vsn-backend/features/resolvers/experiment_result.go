package resolvers

import (
	"context"
	"errors"

	"github.com/seg491X-team36/vsn-backend/domain/model"
)

type ExperimentResultResolvers struct {
}

func (r *ExperimentResultResolvers) User(ctx context.Context, obj *model.ExperimentResult) (model.User, error) {
	return model.User{}, errors.New("not implemented")
}

func (r *ExperimentResultResolvers) Experiment(ctx context.Context, obj *model.ExperimentResult) (model.Experiment, error) {
	return model.Experiment{}, errors.New("not implemented")
}

func (r *ExperimentResultResolvers) Download(ctx context.Context, obj *model.ExperimentResult) (string, error) {
	return "", errors.New("not implemented")
}
