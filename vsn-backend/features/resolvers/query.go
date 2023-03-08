package resolvers

import (
	"context"

	"github.com/google/uuid"
	"github.com/seg491X-team36/vsn-backend/domain/model"
	"github.com/seg491X-team36/vsn-backend/domain/repository"
	"github.com/seg491X-team36/vsn-backend/domain/services"
)

type QueryResolvers struct {
	services.LoginService
	repository.UserRepository
	repository.InviteRepository
	repository.ExperimentRepository
}

func (q *QueryResolvers) Login(ctx context.Context, email string, password string) (*string, error) {
	// call login service
	token, err := q.LoginService.Login(email, password)
	return nullable(token, err), nil
}

func (q *QueryResolvers) User(ctx context.Context, id uuid.UUID) (*model.User, error) {
	// query user by id
	user, err := q.UserRepository.GetUser(ctx, id)
	return nullable(user, err), nil
}

func (q *QueryResolvers) Users(ctx context.Context, state *model.UserAccountState) ([]model.User, error) {
	// query all users by state
	if state == nil {
		*state = model.REGISTERED
	}
	users := q.UserRepository.GetUsersByState(ctx, *state)
	return users, nil
}

func (q *QueryResolvers) Invite(ctx context.Context, id uuid.UUID) (*model.Invite, error) {
	// query invite by id
	invite, err := q.InviteRepository.GetInvite(ctx, id)
	return nullable(invite, err), nil
}

func (q *QueryResolvers) Experiment(ctx context.Context, id uuid.UUID) (*model.Experiment, error) {
	// query experiment by id
	experiment, err := q.ExperimentRepository.GetExperiment(ctx, id)
	return nullable(experiment, err), nil
}

func (q *QueryResolvers) Experiments(ctx context.Context) ([]model.Experiment, error) {
	// query all experiments
	experiments := q.ExperimentRepository.GetExperiments(ctx)
	return experiments, nil
}
