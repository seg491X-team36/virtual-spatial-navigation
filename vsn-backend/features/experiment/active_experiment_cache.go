package experiment

import (
	"errors"

	"github.com/google/uuid"
)

type activeExperimentCache struct {
	experiments map[uuid.UUID]*activeExperiment
}

func (cache *activeExperimentCache) Get(userId uuid.UUID) (*activeExperiment, error) {
	experiment, ok := cache.experiments[userId]
	if !ok {
		return nil, errors.New("not found")
	}
	return experiment, nil
}
