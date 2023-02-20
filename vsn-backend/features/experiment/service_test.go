package experiment

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/seg491X-team36/vsn-backend/domain/model"
	"github.com/stretchr/testify/assert"
)

type inviteRepositoryStub struct {
	Invites []model.Invite
}

func (repository *inviteRepositoryStub) GetPendingInvites(ctx context.Context, userId uuid.UUID) []model.Invite {
	return repository.Invites
}

func TestServicePending(t *testing.T) {
	userId := uuid.New()
	experimentId := uuid.New()

	t.Run("in-progress", func(t *testing.T) {
		service := &Service{
			invites: &inviteRepositoryStub{
				Invites: []model.Invite{
					{ExperimentID: experimentId},
				},
			},
			activeExperiments: &activeExperimentCache{
				experiments: map[uuid.UUID]*activeExperiment{
					userId: {
						UserId:       userId,
						ExperimentId: experimentId,
					},
				},
			},
		}

		res := service.Pending(context.Background(), userId)
		assert.True(t, res.ExperimentInProgress)         // experiment is in progress
		assert.Equal(t, experimentId, *res.ExperimentId) // experiment id is correct
		assert.Equal(t, 0, res.Pending)                  // no pending experiments because the experiment is in progress
	})

	t.Run("not-in-progress-with-invites", func(t *testing.T) {
		// service with no active experiment
		service := &Service{
			invites: &inviteRepositoryStub{
				Invites: []model.Invite{
					{ExperimentID: experimentId},
				},
			},
			activeExperiments: &activeExperimentCache{
				experiments: map[uuid.UUID]*activeExperiment{},
			},
		}

		res := service.Pending(context.Background(), userId)
		assert.False(t, res.ExperimentInProgress)        // experiment not in progress
		assert.Equal(t, experimentId, *res.ExperimentId) // experiment id is correct
		assert.Equal(t, 1, res.Pending)                  // one pending experiment
	})

	t.Run("not-in-progress-without-invites", func(t *testing.T) {
		service := &Service{
			invites: &inviteRepositoryStub{
				Invites: []model.Invite{},
			},
			activeExperiments: &activeExperimentCache{
				experiments: map[uuid.UUID]*activeExperiment{},
			},
		}

		res := service.Pending(context.Background(), userId)
		assert.False(t, res.ExperimentInProgress) // experimennt not in progress
		assert.Nil(t, res.ExperimentId)           // no pending experiment
		assert.Equal(t, 0, res.Pending)           // no pending experiments
	})
}
