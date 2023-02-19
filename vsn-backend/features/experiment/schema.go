package experiment

import (
	"time"

	"github.com/google/uuid"
	"github.com/seg491X-team36/vsn-backend/domain/model"
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
	Experiment *model.ExperimentConfig `json:"experiment"`
	Error      *string                 `json:"error"`
}

type startRoundRequest struct {
	// no request information needed
}

type startRoundResponse struct {
	Status *ExperimentStatus `json:"status"`
	Error  *string           `json:"error"`
}

type stopRoundRequest struct {
	// no request information needed
}

type stopRoundResponse struct {
	Status *ExperimentStatus `json:"status"`
	Error  *string           `json:"error"`
}

type recordDataRequest struct {
	Frames []frame `json:"frames"`
	Events []event `json:"events"`
}

type recordDataResponse struct {
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
