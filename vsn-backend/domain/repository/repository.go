package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/seg491X-team36/vsn-backend/domain/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, email string, source string) (model.User, error)
	GetUser(ctx context.Context, userId uuid.UUID) (model.User, error)
	GetUsersByState(ctx context.Context, state model.UserAccountState) []model.User
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
	UpdateUserState(ctx context.Context, userIds []uuid.UUID, state model.UserAccountState) []model.User
	// get the users who are not invited to an experiment
	GetUsersNotInvited(ctx context.Context, experimentId uuid.UUID) []model.User
}

type InviteRepository interface {
	CreateInvite(ctx context.Context, input model.InviteInput) (model.Invite, error)
	GetInvite(ctx context.Context, inviteId uuid.UUID) (model.Invite, error)
	GetPendingInvitesByExperimentId(ctx context.Context, experimentId uuid.UUID) []model.Invite
	GetPendingInvitesByUserId(ctx context.Context, userId uuid.UUID) []model.Invite
}

type ExperimentRepository interface {
	CreateExperiment(ctx context.Context, input model.ExperimentInput) (model.Experiment, error)
	GetExperiment(ctx context.Context, experimentId uuid.UUID) (model.Experiment, error)
	GetExperiments(ctx context.Context) []model.Experiment // all experiments
}

type ExperimentResultRepository interface {
	CreateExperimentResult(ctx context.Context, input model.ExperimentResultInput) (model.ExperimentResult, error)
	GetExperimentResultsByExperimentId(ctx context.Context, experimentId uuid.UUID) []model.ExperimentResult
	GetExperimentResultsByUserId(ctx context.Context, userId uuid.UUID) []model.ExperimentResult
}
