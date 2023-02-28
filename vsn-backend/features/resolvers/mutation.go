package resolvers

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/seg491X-team36/vsn-backend/codegen/graph"
	"github.com/seg491X-team36/vsn-backend/domain/model"
	"github.com/seg491X-team36/vsn-backend/domain/services"
)

type MutationResolvers struct {
	services.UserService
	services.InviteService
}

func (m *MutationResolvers) UserRegister(ctx context.Context, email string) (graph.UserPayload, error) {
	// register the user
	user, err := payloadWrapper(m.UserService.Register(ctx, email, "MANUAL"))
	return graph.UserPayload{User: user, Error: err}, nil
}

func (m *MutationResolvers) UserSelect(ctx context.Context, input []model.UserSelectInput) ([]graph.UserPayload, error) {
	// select users
	selected := m.UserService.Select(ctx, input)

	// index users by their id
	selectedMap := map[uuid.UUID]model.User{}
	for _, user := range selected {
		selectedMap[user.Id] = user
	}

	users := make([]graph.UserPayload, len(input))

	for i, selectInput := range input {
		user, ok := selectedMap[selectInput.UserID]
		if !ok {
			// user was not updated
			msg := fmt.Errorf("user not found %s", selectInput.UserID).Error()
			users[i] = graph.UserPayload{
				User:  nil,
				Error: &msg,
			}
		} else {
			// user was updated
			users[i] = graph.UserPayload{
				User:  &user,
				Error: nil,
			}
		}
	}

	return users, nil
}

func (m *MutationResolvers) Invite(ctx context.Context, input []model.InviteInput) ([]graph.InvitePayload, error) {
	// results
	invites := make([]graph.InvitePayload, len(input))

	// send invites in parallel
	wg := sync.WaitGroup{}

	for i, req := range input {
		wg.Add(1)
		go func(j int, req model.InviteInput) {
			invite, err := payloadWrapper(m.InviteService.Send(ctx, req))
			invites[j] = graph.InvitePayload{Invite: invite, Error: err}
			wg.Done()
		}(i, req)
	}

	wg.Wait()

	return invites, nil
}
