package experiment

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/seg491X-team36/vsn-backend/domain/model"
)

var (
	errExperimentNotFound           = errors.New("experiment not found")
	errExperimentRoundInProgress    = errors.New("experiment round already in progress")
	errExperimentRoundNotInProgress = errors.New("experiment round not in progress")
)

type verificationEmailRequest struct {
	Email string `json:"email"`
}

type verificationEmailResponse struct {
	// no response information needed. don't leak signed up emails
}

type verificationCodeRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type verificationCodeResponse struct {
	Token string  `json:"token"`
	Error *string `json:"error"`
}

type pendingExperimentsRequest struct {
	// no request information needed.
}

type pendingExperimentsResponse struct {
	ExperimentId         *uuid.UUID `json:"experimentId"`         // experiment id that can be started or resumed
	ExperimentInProgress bool       `json:"experimentInProgress"` // if the experiment is in progress
	Pending              int        `json:"pending"`              // total pending experiments (does not include in progress)
}

type startExperimentRequest struct {
	ExperimentId uuid.UUID `json:"experimentId"` // experiment id to start
}

type startExperimentData struct {
	Experiment model.ExperimentConfig `json:"experiment"`
	Status     model.ExperimentStatus `json:"status"` // always present starting or resuming
	Frame      *frame                 `json:"frame"`  // the last frame recorded. only present when resuming and continuing from last frame
}

type startExperimentResponse struct {
	Data  *startExperimentData `json:"data"`
	Error *string              `json:"error"`
}

type startRoundRequest struct {
	// no request information needed
}

type startRoundResponse struct {
	Status *model.ExperimentStatus `json:"status"`
	Error  *string                 `json:"error"`
}

type stopRoundRequest struct {
	Data experimentData `json:"data"` // the remaining data to ensure all data is recorded before stopping the round
}

type stopRoundResponse struct {
	Status *model.ExperimentStatus `json:"status"`
	Error  *string                 `json:"error"`
}

type recordDataRequest struct {
	Data experimentData `json:"data"`
}

type recordDataResponse struct {
	Error *string `json:"error"`
}

type experimentData struct {
	Frames []frame `json:"frames"`
	Events []event `json:"events"`
}

type frame struct {
	Timestamp time.Time `json:"timestamp"`
	PositionX float64   `json:"x"`
	PositionY float64   `json:"y"`
	PositionZ float64   `json:"z"`
	RotationX float64   `json:"xRot"`
	RotationY float64   `json:"yRot"`
	RotationZ float64   `json:"zRot"`
}

type event struct {
	Name      string    `json:"name"`
	Timestamp time.Time `json:"timestamp"`
}
