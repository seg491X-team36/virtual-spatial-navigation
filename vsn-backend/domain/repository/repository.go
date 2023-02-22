package model

import (
	"context"

	"github.com/google/uuid"
	"github.com/seg491X-team36/vsn-backend/domain/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, email string, source string) (model.User, error)
	GetUser(ctx context.Context, id uuid.UUID) (model.User, error)
	GetUsersByState(ctx context.Context, state model.UserAccountState) ([]model.User, error)
	UpdateUserState(ctx context.Context, state model.UserAccountState, users []uuid.UUID) error
}

type InviteRepository interface {
	CreateInvite(ctx context.Context, input model.InviteInput) (model.Invite, error)
	GetInvite(ctx context.Context, id uuid.UUID) (model.Invite, error)
	GetInvitesByExperimentId(ctx context.Context, experimentId uuid.UUID) []model.Invite
	GetPendingInvites(ctx context.Context, userId uuid.UUID) []model.Invite
}

type ExperimentRepository interface {
}

type ExperimentResultRepository interface {
}
