package model

import (
	"context"

	"github.com/google/uuid"
	"github.com/seg491X-team36/vsn-backend/domain/model"
)

type UserRepository interface {
}

type InviteRepository interface {
	GetInvite(ctx context.Context, id uuid.UUID) (model.Invite, error)
	GetInvitesByExperimentId(ctx context.Context, supervised bool, id uuid.UUID) ([]model.Invite, error)
	CreateInvite(ctx context.Context, input model.InviteInput) (model.Invite, error)
}

type ExperimentRepository interface {
}

type ExperimentResultRepository interface {
}
