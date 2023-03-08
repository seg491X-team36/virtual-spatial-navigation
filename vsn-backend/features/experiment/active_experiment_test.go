package experiment

import (
	"testing"

	"github.com/seg491X-team36/vsn-backend/domain/model"
	"github.com/stretchr/testify/assert"
)

func assertIsReset(t *testing.T, ae *activeExperiment) {
	// relevant fields have been reset
	assert.Nil(t, ae.LatestFrame)
	assert.False(t, ae.RoundInProgress)
	assert.False(t, ae.RewardFound)
}

func assertIsNotReset(t *testing.T, ae *activeExperiment) {
	// used to make sure the round was not reset when resume config is continue
	assert.NotNil(t, ae.LatestFrame)
	assert.True(t, ae.RoundInProgress)
	assert.False(t, ae.RewardFound)
}

func assertEventsRecordedOrdered(t *testing.T, rec *recorderStub, events ...string) {
	for i, evt := range events {
		assert.Equal(t, evt, rec.events[i].Name)
	}
}

func assertEventRecorded(t *testing.T, rec *recorderStub, expected string) {
	recorded := false
	for _, event := range rec.events {
		if event.Name == expected {
			recorded = true
			break
		}
	}
	assert.True(t, recorded)
}

func TestStartRound(t *testing.T) {
	t.Run("in-progress", func(t *testing.T) {
		experiment := &activeExperiment{
			ExperimentStatus: model.ExperimentStatus{
				RoundInProgress: true,
				RoundsCompleted: 0,
				RoundsTotal:     1,
			},
			recorder: &recorderStub{},
		}

		err := experiment.StartRound()
		assert.ErrorIs(t, err, errExperimentRoundInProgress)
	})

	t.Run("not-in-progress", func(t *testing.T) {
		rec := &recorderStub{}
		experiment := &activeExperiment{
			ExperimentStatus: model.ExperimentStatus{
				RoundInProgress: false,
				RoundsCompleted: 0,
				RoundsTotal:     1,
			},
			recorder: rec,
		}

		err := experiment.StartRound()
		assert.NoError(t, err)
		assert.True(t, experiment.RoundInProgress)
		assertEventRecorded(t, rec, eventRoundStart)
	})
}

func TestStopRound(t *testing.T) {
	t.Run("not-in-progress", func(t *testing.T) {
		experiment := &activeExperiment{
			ExperimentStatus: model.ExperimentStatus{
				RoundInProgress: false,
				RoundsCompleted: 0,
				RoundsTotal:     1,
			},
			RewardFound: false,
		}

		err := experiment.StopRound(experimentData{})
		assert.ErrorIs(t, err, errExperimentRoundNotInProgress)
	})

	t.Run("in-progress-reward-not-found", func(t *testing.T) {
		rec := &recorderStub{}
		experiment := &activeExperiment{
			ExperimentStatus: model.ExperimentStatus{
				RoundInProgress: true,
				RoundsCompleted: 0,
				RoundsTotal:     1,
			},
			RewardFound: false,
			recorder:    rec,
			onComplete:  func() {},
		}

		// reward not found
		err := experiment.StopRound(experimentData{})
		assert.ErrorIs(t, err, errExperimentRewardNotFound)

		// reward just found
		err = experiment.StopRound(experimentData{
			Events: []event{{Name: eventRewardFound}},
		})
		assert.NoError(t, err)
		assert.Equal(t, 1, experiment.RoundsCompleted)
		assertIsReset(t, experiment)
		assertEventRecorded(t, rec, eventRoundStop)
	})

	t.Run("in-progress-reward-found", func(t *testing.T) {
		rec := &recorderStub{}
		experiment := &activeExperiment{
			ExperimentStatus: model.ExperimentStatus{
				RoundInProgress: true,
				RoundsCompleted: 0,
				RoundsTotal:     1,
			},
			RewardFound: true,
			recorder:    rec,
			onComplete:  func() {},
		}

		// stop round
		err := experiment.StopRound(experimentData{})
		assert.NoError(t, err)
		assert.Equal(t, 1, experiment.RoundsCompleted)
		assertIsReset(t, experiment)
		assertEventRecorded(t, rec, eventRoundStop)
	})
}

func TestResume(t *testing.T) {
	t.Run("not-in-progress", func(t *testing.T) {
		rec := &recorderStub{}
		experiment := &activeExperiment{
			ExperimentStatus: model.ExperimentStatus{
				RoundInProgress: false,
				RoundsCompleted: 0,
				RoundsTotal:     1,
			},
			recorder: rec,
		}

		experiment.Resume()
		assertEventRecorded(t, rec, eventResumeNoEffect)
	})

	t.Run("in-progress-reward-found", func(t *testing.T) {
		rec := &recorderStub{}
		experiment := &activeExperiment{
			ExperimentStatus: model.ExperimentStatus{
				RoundInProgress: true,
				RoundsCompleted: 0,
				RoundsTotal:     1,
			},
			RewardFound: true,
			LatestFrame: &frame{}, // non nil frame
			recorder:    rec,
			onComplete:  func() {},
		}

		experiment.Resume()
		assert.Equal(t, 1, experiment.RoundsCompleted)   // reset to the next round
		assertIsReset(t, experiment)                     // reset to the next round
		assertEventRecorded(t, rec, eventResumeNoEffect) // recorded as no effect
	})

	t.Run("in-progress-reset", func(t *testing.T) {
		rec := &recorderStub{}
		experiment := &activeExperiment{
			Config: model.ExperimentConfig{
				Resume: model.RESET_ROUND,
			},
			ExperimentStatus: model.ExperimentStatus{
				RoundInProgress: true,
				RoundsCompleted: 0,
				RoundsTotal:     1,
			},
			RewardFound: false,
			LatestFrame: &frame{}, // non nil frame
			recorder:    rec,
		}

		experiment.Resume()
		assert.Equal(t, 0, experiment.RoundsCompleted) // reset to the start of the round
		assertIsReset(t, experiment)                   // reset to the start of the round
		assertEventRecorded(t, rec, eventResumeReset)  // recorded as reset
	})

	t.Run("in-progress-continue", func(t *testing.T) {
		rec := &recorderStub{}
		experiment := &activeExperiment{
			Config: model.ExperimentConfig{
				Resume: model.CONTINUE_ROUND,
			},
			ExperimentStatus: model.ExperimentStatus{
				RoundInProgress: true,
				RoundsCompleted: 0,
				RoundsTotal:     1,
			},
			RewardFound: false,
			LatestFrame: &frame{}, // non nil frame
			recorder:    rec,
			onComplete:  func() {},
		}

		experiment.Resume()
		assert.Equal(t, 0, experiment.RoundsCompleted)   // reset to the next round
		assertIsNotReset(t, experiment)                  // reset to the next round
		assertEventRecorded(t, rec, eventResumeContinue) // recorded as no effect
	})
}

func TestRecord(t *testing.T) {
	rec := &recorderStub{}
	experiment := &activeExperiment{
		RewardFound: false, // IMPORTANT for this test
		LatestFrame: nil,   // IMPORTANT for this test
		recorder:    rec,
	}

	// empty
	experiment.Record(experimentData{})
	assert.Equal(t, 0, len(rec.events))
	assert.Equal(t, 0, len(rec.frames))

	// record some frames and a random event
	experiment.Record(experimentData{
		Frames: []frame{{PositionX: 1}, {PositionX: 2}, {PositionX: 3}},
		Events: []event{{Name: "RANDOM"}},
	})

	assertEventRecorded(t, rec, "RANDOM")
	assert.False(t, experiment.RewardFound)
	assert.Equal(t, 3.0, experiment.LatestFrame.PositionX)

	// record reward found
	experiment.Record(experimentData{
		Events: []event{{Name: eventRewardFound}},
	})
	assert.True(t, experiment.RewardFound)
}
