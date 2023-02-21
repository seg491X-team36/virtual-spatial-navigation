package experiment

import (
	"testing"
	"time"

	"github.com/seg491X-team36/vsn-backend/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestResume(t *testing.T) {
	t.Run("tc1", func(t *testing.T) {
		// not in progress, and RESET_ROUND
		recorder := &recorderStub{}
		experiment := &activeExperiment{
			recorder:    recorder,
			latestFrame: &frame{},
			ExperimentStatus: model.ExperimentStatus{
				RoundInProgress: false, // NOT IN PROGRESS
				RoundNumber:     1,
				RoundsTotal:     2,
			},
			ExperimentConfig: model.ExperimentConfig{
				RoundsTotal:  2,
				ResumeConfig: model.RESET_ROUND, // reset round
			},
		}

		frame := experiment.Resume()

		assert.Nil(t, frame)
		assert.Equal(t, "RESUME:NO_EFFECT", recorder.events[0].Name)
		assert.Equal(t, model.ExperimentStatus{
			RoundInProgress: false,
			RoundNumber:     1,
			RoundsTotal:     2,
		}, experiment.ExperimentStatus)
	})

	t.Run("tc2", func(t *testing.T) {
		// in progress, and RESET_ROUND
		recorder := &recorderStub{}
		experiment := &activeExperiment{
			recorder:    recorder,
			latestFrame: &frame{},
			ExperimentStatus: model.ExperimentStatus{
				RoundInProgress: true,
				RoundNumber:     1,
				RoundsTotal:     2,
			},
			ExperimentConfig: model.ExperimentConfig{
				RoundsTotal:  2,
				ResumeConfig: model.RESET_ROUND, // reset round
			},
		}

		frame := experiment.Resume()

		assert.Nil(t, frame)
		assert.Equal(t, "RESUME:RESET_ROUND", recorder.events[0].Name)
		assert.Equal(t, model.ExperimentStatus{
			RoundInProgress: false,
			RoundNumber:     1,
			RoundsTotal:     2,
		}, experiment.ExperimentStatus)
	})

	t.Run("tc3", func(t *testing.T) {
		// in progress, and CONTINUE_ROUND
		latestFrame := &frame{
			PositionX: 1,
			PositionY: 2,
			PositionZ: 3,
		}
		recorder := &recorderStub{}
		experiment := &activeExperiment{
			recorder:    recorder,
			latestFrame: latestFrame,
			ExperimentStatus: model.ExperimentStatus{
				RoundInProgress: true,
				RoundNumber:     1,
				RoundsTotal:     2,
			},
			ExperimentConfig: model.ExperimentConfig{
				RoundsTotal:  2,
				ResumeConfig: model.CONTINUE_ROUND, // continue round
			},
		}

		frame := experiment.Resume()

		assert.Equal(t, latestFrame, frame)
		assert.Equal(t, "RESUME:CONTINUE_ROUND", recorder.events[0].Name)
		assert.Equal(t, model.ExperimentStatus{
			RoundInProgress: true,
			RoundNumber:     1,
			RoundsTotal:     2,
		}, experiment.ExperimentStatus)
	})

	t.Run("tc4", func(t *testing.T) {
		// not in progress, and CONTINUE_ROUND
		recorder := &recorderStub{}
		experiment := &activeExperiment{
			recorder:    recorder,
			latestFrame: &frame{},
			ExperimentStatus: model.ExperimentStatus{
				RoundInProgress: false,
				RoundNumber:     1,
				RoundsTotal:     2,
			},
			ExperimentConfig: model.ExperimentConfig{
				RoundsTotal:  2,
				ResumeConfig: model.CONTINUE_ROUND, // continue round
			},
		}

		frame := experiment.Resume()

		assert.Nil(t, frame)
		assert.Equal(t, "RESUME:NO_EFFECT", recorder.events[0].Name)
		assert.Equal(t, model.ExperimentStatus{
			RoundInProgress: false,
			RoundNumber:     1,
			RoundsTotal:     2,
		}, experiment.ExperimentStatus)
	})
}

func TestExperimentRecord(t *testing.T) {
	recorder := &recorderStub{}
	experiment := &activeExperiment{
		recorder:    recorder,
		latestFrame: nil,
	}

	experimentEvent := event{Name: "REWARD_FOUND", Timestamp: time.Now().UTC()}

	experiment.Record(experimentData{
		Frames: []frame{{PositionX: 1}, {PositionX: 2}}, // record 2 frames
		Events: []event{experimentEvent},                // record 1 event
	})

	assert.Equal(t, frame{PositionX: 2}, *experiment.latestFrame) // the second frame is the latest frame
	assert.Equal(t, experimentEvent, recorder.events[0])          // the event was recorded
}
