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

func recorderStubFactory(stub *recorderStub) recorderFactory {
	return func(params recorderParams) recorder {
		return stub
	}
}

func (r *recorderStub) Record(round int, data experimentData) {
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

type experimentResultRepositoryStub struct {
	input model.ExperimentResultInput
}

func (repository *experimentResultRepositoryStub) CreateExperimentResult(ctx context.Context, input model.ExperimentResultInput) (model.ExperimentResult, error) {
	repository.input = input
	return model.ExperimentResult{}, nil
}

func TestServicePending(t *testing.T) {
	userId := uuid.New()
	experimentId := uuid.New()

	t.Run("active-one-invite", func(t *testing.T) {
		service := &Service{
			invites: &inviteRepositoryStub{
				Invites: []model.Invite{
					{ExperimentID: experimentId},
				},
			},
			activeExperiments: &activeExperimentCache{
				experiments: map[uuid.UUID]*activeExperiment{
					userId: {
						ExperimentId: experimentId,
						recorder:     &recorderStub{},
					},
				},
			},
		}

		res := service.Pending(context.Background(), userId)
		assert.True(t, res.ExperimentInProgress)
		assert.Equal(t, experimentId, *res.ExperimentId) // experiment id is correct
		assert.Equal(t, 0, res.Pending)                  // no pending experiments because the experiment is in progress
	})

	t.Run("no-active-one-invite", func(t *testing.T) {
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

		assert.Equal(t, pendingExperimentsResponse{
			ExperimentId:         &experimentId,
			ExperimentInProgress: false,
			Pending:              1,
		}, res)
	})

	t.Run("no-active-no-invite", func(t *testing.T) {
		service := &Service{
			invites: &inviteRepositoryStub{
				Invites: []model.Invite{},
			},
			activeExperiments: &activeExperimentCache{
				experiments: map[uuid.UUID]*activeExperiment{},
			},
		}

		res := service.Pending(context.Background(), userId)
		assert.Equal(t, pendingExperimentsResponse{
			ExperimentId:         nil,
			ExperimentInProgress: false,
			Pending:              0,
		}, res)
	})
}

func TestServiceStartAndStopRound(t *testing.T) {
	experimentId := uuid.New()
	userId := uuid.New()
	trackingId := uuid.New()

	results := &experimentResultRepositoryStub{}

	service := &Service{
		invites:           &inviteRepositoryStub{},
		experiments:       &experimentRepositoryStub{},
		experimentResults: results,
		activeExperiments: &activeExperimentCache{
			experiments: map[uuid.UUID]*activeExperiment{},
		},
	}

	ae := &activeExperiment{
		ExperimentId: experimentId,
		ExperimentStatus: model.ExperimentStatus{
			RoundInProgress: false,
			RoundsCompleted: 0,
			RoundsTotal:     1,
		},
		recorder:    &recorderStub{},
		LatestFrame: &frame{},
		onComplete:  service.onComplete(experimentId, trackingId, userId),
	}

	service.activeExperiments.Set(userId, ae)

	ctx := context.Background()

	// start round for a user with no active experiments
	status, err := service.StartRound(ctx, uuid.New())
	assert.ErrorIs(t, err, errExperimentNotFound)
	assert.Nil(t, status)

	// stop round for a user with no active experiments
	status, err = service.StopRound(ctx, uuid.New(), experimentData{})
	assert.ErrorIs(t, err, errExperimentNotFound)
	assert.Nil(t, status)

	// start round
	res, err := service.StartRound(ctx, userId)
	assert.NoError(t, err)
	assert.Equal(t, &model.ExperimentStatus{
		RoundInProgress: true,
		RoundsCompleted: 0,
		RoundsTotal:     1,
	}, res)

	ae.RewardFound = true

	// stop round
	res, err = service.StopRound(ctx, userId, experimentData{})
	assert.NoError(t, err)
	assert.Equal(t, &model.ExperimentStatus{
		RoundInProgress: false,
		RoundsCompleted: 1,
		RoundsTotal:     1,
	}, res)

	// verifying onComplete was executed
	_, err = service.activeExperiments.Get(userId)
	assert.ErrorIs(t, err, errExperimentNotFound)

	assert.Equal(t, model.ExperimentResultInput{
		TrackingId:   trackingId,
		UserId:       userId,
		ExperimentId: experimentId,
	}, results.input)
}

func TestServiceStartExperiment(t *testing.T) {
	userId := uuid.New()
	experimentId := uuid.New()

	invite := model.Invite{
		ExperimentID: experimentId,
	}

	experiment := model.Experiment{
		Id: experimentId,
		Config: model.ExperimentConfig{
			RoundsTotal: 2,
			Resume:      model.RESET_ROUND,
		},
	}

	ctx := context.Background()

	t.Run("start-experiment-happy-path", func(t *testing.T) {
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
			recorderFactory: recorderStubFactory(&recorderStub{}),
		}

		res, err := service.StartExperiment(ctx, userId, experimentId)

		assert.NoError(t, err)
		assert.Equal(t, experiment.Config, res.Config)
		assert.Equal(t, model.NewExperimentStatus(experiment.Config.RoundsTotal), res.Status)
	})

	t.Run("resume-experiment-happy-path", func(t *testing.T) {
		// test resume with experiment config reset to resume
		ae := &activeExperiment{
			ExperimentId: experimentId,
			ExperimentStatus: model.ExperimentStatus{
				RoundInProgress: true, // round in progress
				RoundsCompleted: 1,
				RoundsTotal:     2,
			},
			Config:      experiment.Config,
			RewardFound: false,    // reward not found
			LatestFrame: &frame{}, // round in progress, last frame is not nil
			recorder:    &recorderStub{},
		}

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
					userId: ae,
				},
			},
		}

		ctx := context.Background()
		res, err := service.StartExperiment(ctx, userId, experimentId)
		assert.NoError(t, err)
		assert.Equal(t, model.ExperimentStatus{
			RoundInProgress: false,
			RoundsCompleted: 1,
			RoundsTotal:     2,
		}, res.Status)
		assert.Nil(t, res.Frame)
		assertIsReset(t, ae)

	})

	t.Run("start-no-experiments", func(t *testing.T) {
		service := &Service{
			invites: &inviteRepositoryStub{
				Invites: []model.Invite{}, // empty
			},
			experiments: &experimentRepositoryStub{
				Experiment: experiment,
				Err:        nil,
			},
			activeExperiments: &activeExperimentCache{
				experiments: map[uuid.UUID]*activeExperiment{}, // empty
			},
		}

		_, err := service.StartExperiment(ctx, userId, experimentId)
		assert.ErrorIs(t, err, errExperimentNotFound)
	})

	t.Run("start-incorrect-experiment-id", func(t *testing.T) {
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
