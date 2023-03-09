package experiment

import (
	"time"

	"github.com/google/uuid"
	"github.com/seg491X-team36/vsn-backend/domain/model"
)

const (
	eventRoundStart     = "ROUND:START"
	eventRoundStop      = "ROUND:STOP"
	eventResumeNoEffect = "RESUME:NO_EFFECT"
	eventResumeContinue = "RESUME:CONTINUE_ROUND"
	eventResumeReset    = "RESUME:RESET_ROUND"
	eventRewardFound    = "ROUND:REWARD_FOUND"
)

type activeExperiment struct {
	// experiment id, config
	ExperimentId uuid.UUID
	Config       model.ExperimentConfig

	// experiment status, and round information fields
	model.ExperimentStatus
	LatestFrame *frame
	RewardFound bool

	// extension logic
	recorder   recorder
	onComplete func()
}

func (ae *activeExperiment) Status() model.ExperimentStatus {
	return ae.ExperimentStatus
}

func (ae *activeExperiment) StartExperimentData() *startExperimentData {
	return &startExperimentData{
		Config: ae.Config,
		Status: ae.ExperimentStatus,
		Frame:  ae.LatestFrame,
	}
}

func (ae *activeExperiment) StartRound() error {
	// can't start the round if the round is in progress
	if ae.RoundInProgress {
		return errExperimentRoundInProgress
	}
	ae.RecordEvent(eventRoundStart)

	// update the round in progress to true
	ae.RoundInProgress = true
	return nil
}

func (ae *activeExperiment) StopRound(data experimentData) error {
	// can't stop the round if the round is not in progress
	if !ae.RoundInProgress {
		return errExperimentRoundNotInProgress
	}

	// record data in case the REWARD_FOUND event and to save remaining frames
	ae.Record(data)

	// can't stop the round if the reward has not been found
	if !ae.RewardFound {
		return errExperimentRewardNotFound
	}
	ae.RecordEvent(eventRoundStop)

	// update the number of rounds completed
	ae.RoundsCompleted += 1
	ae.reset()

	if ae.Complete() {
		ae.onComplete()
	}

	return nil
}

func (ae *activeExperiment) Resume() {
	// if the round is in progress, then there's nothing to resume
	if !ae.RoundInProgress {
		ae.RecordEvent(eventResumeNoEffect)
		return
	}

	// if the reward has been found, then stop the round
	if ae.RewardFound {
		_ = ae.StopRound(experimentData{})
		ae.RecordEvent(eventResumeNoEffect) // record the event as part of the next round
		return
	}

	// resume and continue round
	if ae.Config.Resume == model.CONTINUE_ROUND && ae.LatestFrame != nil {
		ae.RecordEvent(eventResumeContinue)
		return
	}

	// resume and reset round
	ae.RecordEvent(eventResumeReset)
	ae.reset()
}

func (ae *activeExperiment) Record(data experimentData) {
	// update the latest frame
	if n := len(data.Frames); n > 0 {
		ae.LatestFrame = &data.Frames[n-1]
	}

	// update the reward was found
	for _, event := range data.Events {
		if event.Name == eventRewardFound {
			ae.RewardFound = true
			break
		}
	}

	ae.recorder.Record(ae.RoundsCompleted, data)
}

func (ae *activeExperiment) RecordEvent(name string) {
	ae.recorder.Record(ae.RoundsCompleted, experimentData{
		Frames: []frame{},
		Events: []event{{Name: name, Timestamp: time.Now().UTC()}},
	})
}

// reset round information
func (ae *activeExperiment) reset() {
	ae.RoundInProgress = false // reset round in progress
	ae.LatestFrame = nil       // reset latest frame
	ae.RewardFound = false     // reset the reward
}
