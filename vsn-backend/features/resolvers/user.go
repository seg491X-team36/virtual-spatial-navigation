package resolvers

import (
	"context"

	"github.com/seg491X-team36/vsn-backend/domain/model"
	"github.com/seg491X-team36/vsn-backend/domain/repository"
)

type UserResolvers struct {
	repository.InviteRepository
	repository.ExperimentResultRepository
}

func (r *UserResolvers) Invites(ctx context.Context, user *model.User) ([]model.Invite, error) {
	// call the invite repository
	invites := r.InviteRepository.GetPendingInvitesByUserId(ctx, user.Id)
	return invites, nil
}

func (r *UserResolvers) Results(ctx context.Context, user *model.User) ([]model.ExperimentResult, error) {
	// call the experiment repository
	results := r.ExperimentResultRepository.GetExperimentResultsByUserId(ctx, user.Id)
	return results, nil
}
