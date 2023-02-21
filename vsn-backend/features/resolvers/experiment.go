package resolvers

import (
	"context"
	"errors"

	"github.com/seg491X-team36/vsn-backend/domain/model"
)

type ExperimentResolvers struct {
}

func (r *ExperimentResolvers) Arena(ctx context.Context, obj *model.Experiment) (model.Arena, error) {
	return model.Arena{}, errors.New("not implemented")
}

func (r *ExperimentResolvers) Results(ctx context.Context, obj *model.Experiment) ([]model.ExperimentResult, error) {
	return []model.ExperimentResult{}, errors.New("not implemented")
}
