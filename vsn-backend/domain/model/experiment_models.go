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
	RoundsTotal    int                    `json:"roundsTotal"`
	ResumeConfig   ExperimentResumeConfig `json:"resumeConfig"`
	SpawnSequence  []int                  `json:"spawnSequence"`
	RewardPosition int                    `json:"rewardPosition"`
	Arena          Arena                  `json:"arena"`
}

type Arena struct {
	Objects         []string   `json:"objects"` // string identifier known by unity game
	RewardPositions []Position `json:"rewardPositions"`
	SpawnPositions  []Position `json:"spawnPositions"`
}

type ArenaObject struct {
	Object   string   `json:"object"` // string identifier
	Position Position `json:"position"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
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
	TrackingId   uuid.UUID // "TrackingId" in experiments package
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
	RoundsCompleted int  `json:"roundsCompleted"`
	RoundsTotal     int  `json:"-"` // omitted in json
}

func NewExperimentStatus(roundsTotal int) ExperimentStatus {
	return ExperimentStatus{
		RoundInProgress: false,
		RoundsCompleted: 0,
		RoundsTotal:     roundsTotal,
	}
}

func (s ExperimentStatus) Complete() bool {
	return s.RoundsCompleted == s.RoundsTotal && !s.RoundInProgress
}
