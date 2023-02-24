package experiment

import (
	"time"

	"github.com/google/uuid"
	"github.com/seg491X-team36/vsn-backend/domain/model"
)

type activeExperiment struct {
	TrackingId   uuid.UUID // id used to save the results
	ExperimentId uuid.UUID
	UserId       uuid.UUID
	recorder     recorder
	latestFrame  *frame // track the latest frame
	model.ExperimentStatus
	model.ExperimentConfig
}

func (ae *activeExperiment) Resume() *frame {
	/* case #1 the round must be reset and it is not in progress
	- nothing needs to be reset */
	if ae.ResumeConfig == model.RESET_ROUND && !ae.RoundInProgress {
		ae.RecordEvent("RESUME:NO_EFFECT")
		return nil
	}

	// case #2 the round must be reset and it is in progress
	if ae.ResumeConfig == model.RESET_ROUND && ae.RoundInProgress {
		ae.RecordEvent("RESUME:RESET_ROUND")
		ae.RoundInProgress = false
		return nil
	}

	// case #3 the round must be continued and it is in progress
	if ae.ResumeConfig == model.CONTINUE_ROUND && ae.RoundInProgress {
		ae.RecordEvent("RESUME:CONTINUE_ROUND")
		return ae.latestFrame
	}

	/* case #4 the round must be continued but it is not in progress
	- no frame can be returned */
	ae.RecordEvent("RESUME:NO_EFFECT")
	return nil
}

func (ae *activeExperiment) StartRound() (model.ExperimentStatus, error) {
	if ae.RoundInProgress {
		return ae.ExperimentStatus, errExperimentRoundInProgress
	}

	// update round in progress
	ae.RecordEvent("ROUND:START")
	ae.RoundInProgress = true
	return ae.ExperimentStatus, nil
}

func (ae *activeExperiment) StopRound(data experimentData) (model.ExperimentStatus, error) {
	if !ae.RoundInProgress {
		return ae.ExperimentStatus, errExperimentRoundNotInProgress
	}
	ae.Record(data) // record data

	// update round in progress and go to the next round
	ae.RecordEvent("ROUND:STOP")
	ae.RoundInProgress = false
	ae.RoundsCompleted += 1
	ae.latestFrame = nil // reset the latest frame
	return ae.ExperimentStatus, nil
}

func (ae *activeExperiment) Record(data experimentData) {
	if n := len(data.Frames); n > 0 {
		ae.latestFrame = &data.Frames[n-1] // update the latest frame
	}
	ae.recorder.Record(ae.RoundsCompleted, data)
}

func (ae *activeExperiment) RecordEvent(name string) {
	ae.Record(experimentData{
		Frames: []frame{},
		Events: []event{{Name: name, Timestamp: time.Now().UTC()}},
	})
}
