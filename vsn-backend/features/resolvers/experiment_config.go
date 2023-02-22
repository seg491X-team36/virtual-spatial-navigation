package resolvers

import (
	"context"
	"errors"

	"github.com/seg491X-team36/vsn-backend/domain/model"
)

type ExperimentConfigResolvers struct {
}

func (r *ExperimentConfigResolvers) Rounds(ctx context.Context, obj *model.ExperimentConfig) (int, error) {
	return 0, errors.New("not implemented")
}

func (r *ExperimentConfigResolvers) Resume(ctx context.Context, obj *model.ExperimentConfig) (model.ExperimentResumeConfig, error) {
	return model.CONTINUE_ROUND, errors.New("not implemented")
}
