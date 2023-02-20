package experiment

import (
	"github.com/google/uuid"
	"github.com/seg491X-team36/vsn-backend/domain/model"
)

type activeExperiment struct {
	TrackingId   uuid.UUID // id used to save the results
	ExperimentId uuid.UUID
	UserId       uuid.UUID
	model.ExperimentStatus
	model.ExperimentConfig
}

func (ae *activeExperiment) StartRound() (model.ExperimentStatus, error) {
	if ae.RoundInProgress {
		return ae.ExperimentStatus, errExperimentRoundInProgress
	}

	// update round in progress and round number
	ae.RoundInProgress = true
	ae.RoundNumber += 1
	return ae.ExperimentStatus, nil
}

func (ae *activeExperiment) StopRound(data experimentData) (model.ExperimentStatus, error) {
	if !ae.RoundInProgress {
		return ae.ExperimentStatus, errExperimentRoundNotInProgress
	}
	ae.Record(data) // record data

	// update round in progress
	ae.RoundInProgress = false
	return ae.ExperimentStatus, nil
}

func (ae *activeExperiment) Record(data experimentData) {
}
