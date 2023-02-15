package experiment

import (
	"time"

	"github.com/google/uuid"
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
	ExperimentId *uuid.UUID `json:"experimentId"` // the first pending experiment id
	Pending      int        `json:"pending"`      // total pending experiments
}

type startExperimentRequest struct {
	ExperimentId uuid.UUID `json:"experimentId"`
}

type startExperimentResponse struct {
	Experiment *experimentConfig `json:"experiment"`
	Error      *string           `json:"error"`
}

type startRoundRequest struct {
	// no request information needed
}

type startRoundResponse struct {
	Status experimentStatus `json:"status"`
	Error  *string          `json:"error"`
}

type stopRoundRequest struct {
	// no request information needed
}

type stopRoundResponse struct {
	Status experimentStatus `json:"status"`
	Error  *string          `json:"error"`
}

type frameRequest struct {
	Frames []frame `json:"frames"`
	Events []event `json:"events"`
}

type frameResponse struct {
	Error *string `json:"error"`
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

type experimentConfig struct {
	// TODO
	// arena id
	// doors
	// objects
	// reward position
}

type experimentStatus struct {
	RoundNumber     int `json:"round"`
	RoundsRemaining int `json:"remaining"`
	RoundsTotal     int `json:"total"`
}
