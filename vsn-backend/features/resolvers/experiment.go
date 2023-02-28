package resolvers

import (
	"context"

	"github.com/seg491X-team36/vsn-backend/domain/model"
	"github.com/seg491X-team36/vsn-backend/domain/repository"
)

type ExperimentResolvers struct {
	repository.ExperimentResultRepository
	repository.InviteRepository
	repository.UserRepository
}

func (r *ExperimentResolvers) Results(ctx context.Context, experiment *model.Experiment) ([]model.ExperimentResult, error) {
	// call the experiment result repository
	results := r.ExperimentResultRepository.GetExperimentResultsByExperimentId(ctx, experiment.Id)
	return results, nil
}

func (r *ExperimentResolvers) Pending(ctx context.Context, experiment *model.Experiment) ([]model.Invite, error) {
	// call the invites repository
	invites := r.InviteRepository.GetPendingInvitesByExperimentId(ctx, experiment.Id)
	return invites, nil
}

func (r *ExperimentResolvers) UsersNotInvited(ctx context.Context, experiment *model.Experiment) ([]model.User, error) {
	// call the users repository
	users := r.UserRepository.GetUsersNotInvited(ctx, experiment.Id)
	return users, nil
}
