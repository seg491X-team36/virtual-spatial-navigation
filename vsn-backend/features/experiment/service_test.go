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

func TestStartAndStopRound(t *testing.T) {
	userId1 := uuid.New()
	userId2 := uuid.New()

	experiments := &activeExperimentCache{
		experiments: map[uuid.UUID]*activeExperiment{
			userId1: {
				UserId: userId1,
				ExperimentStatus: model.ExperimentStatus{
					RoundInProgress: false,
					RoundNumber:     0,
					RoundsTotal:     2,
				},
			},
		},
	}

	service := &Service{
		invites:           &inviteRepositoryStub{},
		activeExperiments: experiments,
	}

	ctx := context.Background()

	// start round for a user with no active experiments
	status, err := service.StartRound(ctx, userId2)
	assert.ErrorIs(t, err, errExperimentNotFound)
	assert.Nil(t, status)

	// stop round for a user with no active experiments
	status, err = service.StopRound(ctx, userId2, experimentData{})
	assert.ErrorIs(t, err, errExperimentNotFound)
	assert.Nil(t, status)

	// stop round for a user with round not in progress
	status, err = service.StopRound(ctx, userId1, experimentData{})
	assert.ErrorIs(t, err, errExperimentRoundNotInProgress)
	assert.Equal(t, &model.ExperimentStatus{
		RoundInProgress: false,
		RoundNumber:     0,
		RoundsTotal:     2,
	}, status) // still returns the status even if the round is not stopped

	for i := 1; i <= 2; i++ {
		// start round
		status, err = service.StartRound(ctx, userId1)
		expected := &model.ExperimentStatus{
			RoundInProgress: true,
			RoundNumber:     i,
			RoundsTotal:     2,
		}
		assert.NoError(t, err)
		assert.Equal(t, expected, status)

		// start round gives an error because the round is in progress
		status, err = service.StartRound(ctx, userId1)
		assert.Error(t, errExperimentRoundInProgress)
		assert.Equal(t, expected, status) // status does not change

		// stop round
		status, err = service.StopRound(ctx, userId1, experimentData{})
		assert.NoError(t, err)
		assert.Equal(t, &model.ExperimentStatus{
			RoundInProgress: false,
			RoundNumber:     i,
			RoundsTotal:     2,
		}, status)
	}

	_, err = experiments.Get(userId1)
	assert.Error(t, errExperimentNotFound) // the experiment was deleted

}
