package resolvers

import (
	"context"

	"github.com/google/uuid"
	"github.com/seg491X-team36/vsn-backend/domain/model"
	"github.com/seg491X-team36/vsn-backend/features/dataloader"
)

type ExperimentResultResolvers struct {
	Users       dataloader.Loader[uuid.UUID, model.User]
	Experiments dataloader.Loader[uuid.UUID, model.Experiment]
}

func (r *ExperimentResultResolvers) User(ctx context.Context, result *model.ExperimentResult) (model.User, error) {
	return r.Users.Get(ctx, result.UserId)
}

func (r *ExperimentResultResolvers) Experiment(ctx context.Context, result *model.ExperimentResult) (model.Experiment, error) {
	return r.Experiments.Get(ctx, result.ExperimentId)
}
