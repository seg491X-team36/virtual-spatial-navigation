package model

import (
	"time"

	"github.com/google/uuid"
)

type ExperimentResumeConfig string

const (
	CONTINUE_ROUND = ExperimentResumeConfig("CONTINUE_ROUND")
	RESET_ROUND    = ExperimentResumeConfig("RESET_ROUND")
)

type Experiment struct {
	Id          uuid.UUID
	Name        string
	Description string
	Config      ExperimentConfig
}

type ExperimentConfig struct {
	RoundsTotal  int
	ResumeConfig ExperimentResumeConfig
}

type ExperimentResult struct {
	Id           uuid.UUID
	CreatedAt    time.Time
	UserId       uuid.UUID
	ExperimentId uuid.UUID
}

type ExperimentInput struct {
	Name        string
	Description string
	ExperimentConfig
}

type ExperimentResultInput struct {
	Id           uuid.UUID // "TrackingId" in experiments package
	UserId       uuid.UUID
	ExperimentId uuid.UUID
}

/* experiment ExperimentStatus struct
if an experiment has 3 rounds the state should go:

start experiment -> {"roundInProgress": false, "roundNumber": 0}
start round -> {"roundInProgress": true, "roundNumber": 0}
stop round -> {"roundInProgress": false, "roundNumber": 1}
start round -> {"roundInProgress": true "roundNumber": 1}
stop round -> {"roundInProgress": false, "roundNumber": 2}
start round -> {"roundInProgress": true, "roundNumber": 2}
stop round -> {"roundInProgress": false, "roundNumber": 3} DONE
*/
type ExperimentStatus struct {
	RoundInProgress bool `json:"roundInProgress"`
	RoundNumber     int  `json:"roundNumber"`
	RoundsTotal     int  `json:"roundsTotal"`
}

func (s ExperimentStatus) Done() bool {
	return s.RoundNumber == s.RoundsTotal && !s.RoundInProgress
}
