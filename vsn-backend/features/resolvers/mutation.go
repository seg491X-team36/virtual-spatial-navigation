package resolvers

import (
	"context"
	"errors"

	"github.com/seg491X-team36/vsn-backend/codegen/graph"
	"github.com/seg491X-team36/vsn-backend/domain/model"
)

type MutationResolvers struct {
}

func (m *MutationResolvers) UserRegister(ctx context.Context, email string) (graph.UserPayload, error) {
	return graph.UserPayload{}, errors.New("not implemented")
}

func (m *MutationResolvers) UserSelect(ctx context.Context, input []model.UserSelectionInput) ([]graph.UserSelectionPayload, error) {
	return []graph.UserSelectionPayload{}, errors.New("not implemented")
}

func (m *MutationResolvers) Invite(ctx context.Context, input []model.InviteInput) ([]graph.InvitePayload, error) {
	return []graph.InvitePayload{}, errors.New("not implemented")
}
