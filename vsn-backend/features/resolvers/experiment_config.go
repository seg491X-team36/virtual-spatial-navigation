package resolvers

import (
	"context"

	"github.com/seg491X-team36/vsn-backend/domain/model"
)

type ExperimentConfigResolvers struct {
}

func (r *ExperimentConfigResolvers) Rounds(ctx context.Context, config *model.ExperimentConfig) (int, error) {
	return config.RoundsTotal, nil
}

func (r *ExperimentConfigResolvers) Resume(ctx context.Context, config *model.ExperimentConfig) (model.ExperimentResumeConfig, error) {
	return config.Resume, nil
}
