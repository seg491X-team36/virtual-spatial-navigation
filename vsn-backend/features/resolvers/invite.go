package resolvers

import (
	"context"

	"github.com/google/uuid"
	"github.com/seg491X-team36/vsn-backend/domain/model"
	"github.com/seg491X-team36/vsn-backend/features/dataloader"
)

type InviteResolvers struct {
	Users       dataloader.Loader[uuid.UUID, model.User]
	Experiments dataloader.Loader[uuid.UUID, model.Experiment]
}

func (r *InviteResolvers) User(ctx context.Context, invite *model.Invite) (model.User, error) {
	// call the users data loader
	return r.Users.Get(ctx, invite.UserID)
}

func (r *InviteResolvers) Experiment(ctx context.Context, invite *model.Invite) (model.Experiment, error) {
	// call the experiments dataloader
	return r.Experiments.Get(ctx, invite.UserID)
}
