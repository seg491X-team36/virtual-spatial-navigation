package experiment

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/seg491X-team36/vsn-backend/domain/model"
	"github.com/stretchr/testify/assert"
)

type recorderStub struct {
	events []event
	frames []frame
}

func (r *recorderStub) Record(metadata recorderMetadata, data experimentData) {
	r.events = append(r.events, data.Events...)
	r.frames = append(r.frames, data.Frames...)
}

type inviteRepositoryStub struct {
	inviteRepository
	Invites []model.Invite
}

func (repository *inviteRepositoryStub) GetPendingInvites(ctx context.Context, userId uuid.UUID) []model.Invite {
	return repository.Invites
}

type experimentRepositoryStub struct {
	experimentRepository
	Experiment model.Experiment
	Err        error
}

func (repository *experimentRepositoryStub) GetExperiment(ctx context.Context, experimentId uuid.UUID) (model.Experiment, error) {
	return repository.Experiment, repository.Err
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
						recorder:     &recorderStub{},
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

func TestServiceStartAndStopRound(t *testing.T) {
	userId1 := uuid.New()
	userId2 := uuid.New()

	experiment := &activeExperiment{
		UserId: userId1,
		ExperimentStatus: model.ExperimentStatus{
			RoundInProgress: false,
			RoundNumber:     0,
			RoundsTotal:     2,
		},
		recorder:    &recorderStub{},
		latestFrame: &frame{},
	}

	experiments := &activeExperimentCache{
		experiments: map[uuid.UUID]*activeExperiment{
			userId1: experiment,
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

	for i := 0; i < 2; i++ {
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
		assert.ErrorIs(t, err, errExperimentRoundInProgress)
		assert.Equal(t, expected, status) // status does not change

		// stop round
		status, err = service.StopRound(ctx, userId1, experimentData{})
		assert.NoError(t, err)
		assert.Equal(t, &model.ExperimentStatus{
			RoundInProgress: false,
			RoundNumber:     i + 1,
			RoundsTotal:     2,
		}, status)

		assert.Nil(t, experiment.latestFrame) // the latest frame was reset
	}

	_, err = service.StartRound(ctx, userId1)
	assert.ErrorIs(t, err, errExperimentNotFound) // the experiment was deleted
}

func TestServiceStartExperiment(t *testing.T) {
	userId := uuid.New()
	experimentId := uuid.New()

	invite := model.Invite{
		ID:           uuid.New(),
		UserID:       userId,
		ExperimentID: experimentId,
	}

	experiment := model.Experiment{
		Id: experimentId,
		Config: model.ExperimentConfig{
			RoundsTotal:  2,
			ResumeConfig: model.RESET_ROUND,
		},
	}

	ctx := context.Background()

	t.Run("happy-path", func(t *testing.T) {
		service := &Service{
			invites: &inviteRepositoryStub{
				Invites: []model.Invite{invite},
			},
			experiments: &experimentRepositoryStub{
				Experiment: experiment,
				Err:        nil,
			},
			activeExperiments: &activeExperimentCache{
				experiments: map[uuid.UUID]*activeExperiment{},
			},
		}

		res, err := service.StartExperiment(ctx, userId, experimentId)

		// return values all make sense
		assert.NoError(t, err)
		assert.Nil(t, res.Frame)
		assert.Equal(t, experiment.Config, res.Experiment)
		assert.Equal(t, model.ExperimentStatus{
			RoundInProgress: false,
			RoundNumber:     0,
			RoundsTotal:     2,
		}, res.Status)
	})

	t.Run("resume-experiment", func(t *testing.T) {
		service := &Service{
			invites: &inviteRepositoryStub{
				Invites: []model.Invite{invite},
			},
			experiments: &experimentRepositoryStub{
				Experiment: experiment,
				Err:        nil,
			},
			activeExperiments: &activeExperimentCache{
				experiments: map[uuid.UUID]*activeExperiment{
					userId: {
						recorder:     &recorderStub{},
						latestFrame:  nil,
						TrackingId:   uuid.New(),
						ExperimentId: experimentId,
						UserId:       userId,
						ExperimentStatus: model.ExperimentStatus{
							RoundInProgress: false,
							RoundNumber:     1,
							RoundsTotal:     2,
						},
						ExperimentConfig: experiment.Config,
					},
				},
			},
		}

		ctx := context.Background()
		res, err := service.StartExperiment(ctx, userId, experimentId)
		assert.NoError(t, err)
		assert.Equal(t, model.ExperimentStatus{
			RoundInProgress: false,
			RoundNumber:     1,
			RoundsTotal:     2,
		}, res.Status)
	})

	t.Run("no-pending-experiments", func(t *testing.T) {
		service := &Service{
			invites: &inviteRepositoryStub{
				Invites: []model.Invite{},
			},
			experiments: &experimentRepositoryStub{
				Experiment: experiment,
				Err:        nil,
			},
			activeExperiments: &activeExperimentCache{
				experiments: map[uuid.UUID]*activeExperiment{},
			},
		}

		_, err := service.StartExperiment(ctx, userId, experimentId)
		assert.ErrorIs(t, err, errExperimentNotFound)
	})

	t.Run("pending-experiment-id-does-not-match", func(t *testing.T) {
		service := &Service{
			invites: &inviteRepositoryStub{
				Invites: []model.Invite{invite},
			},
			experiments: &experimentRepositoryStub{
				Experiment: experiment,
				Err:        nil,
			},
			activeExperiments: &activeExperimentCache{
				experiments: map[uuid.UUID]*activeExperiment{},
			},
		}

		newExperimentId := uuid.New()

		_, err := service.StartExperiment(ctx, userId, newExperimentId)
		assert.ErrorIs(t, err, errExperimentNotFound)
	})
}

func TestServiceRecord(t *testing.T) {
	userId := uuid.New()
	service := &Service{
		activeExperiments: &activeExperimentCache{
			experiments: map[uuid.UUID]*activeExperiment{
				userId: {
					recorder: &recorderStub{},
				},
			},
		},
	}

	// user with active experiment
	err := service.Record(context.Background(), userId, experimentData{})
	assert.NoError(t, err)

	// user without active experiment
	newUserId := uuid.New()
	err = service.Record(context.Background(), newUserId, experimentData{})
	assert.ErrorIs(t, err, errExperimentNotFound)
}
