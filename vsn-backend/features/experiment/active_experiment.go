package experiment

import (
	"github.com/google/uuid"
)

type activeExperiment struct {
	TrackingId   uuid.UUID // id used to save the results
	ExperimentId uuid.UUID
	UserId       uuid.UUID
}
