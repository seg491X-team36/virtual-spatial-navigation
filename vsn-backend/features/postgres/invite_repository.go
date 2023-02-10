package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/seg491X-team36/vsn-backend/codegen/db"
	"github.com/seg491X-team36/vsn-backend/domain/model"
)

func convertInvites(records []db.Invite) []model.Invite {
	invites := make([]model.Invite, len(records))
	for i, record := range records {
		invites[i] = model.Invite(record)
	}
	return invites
}

type InviteRepository struct {
	Query *db.Queries
}

// Gets an invite by id
func (repository *InviteRepository) GetInvite(
	ctx context.Context,
	id uuid.UUID,
) (model.Invite, error) {
	invite, err := repository.Query.GetInvite(ctx, id)
	return model.Invite(invite), err
}

// Gets all the invites for a list of experiment ids
func (repository *InviteRepository) GetInvitesByExperimentId(
	ctx context.Context,
	supervised bool,
	id uuid.UUID,
) ([]model.Invite, error) {
	invites, err := repository.Query.GetInvitesByExperimentId(ctx, db.GetInvitesByExperimentIdParams{
		Supervised:   supervised,
		ExperimentID: id,
	})
	return convertInvites(invites), err
}

// Creates an invite
func (repository *InviteRepository) CreateInvite(
	ctx context.Context,
	input model.InviteInput,
) (model.Invite, error) {
	userId, err := uuid.Parse(input.UserID)
	if err != nil {
		return model.Invite{}, err
	}
	expId, err := uuid.Parse(input.ExperimentID)
	if err != nil {
		return model.Invite{}, err
	}
	invite, err := repository.Query.CreateInvite(ctx, db.CreateInviteParams{
		UserID:       userId,
		ExperimentID: expId,
		Supervised:   input.Supervised,
	})
	return model.Invite(invite), err
}
