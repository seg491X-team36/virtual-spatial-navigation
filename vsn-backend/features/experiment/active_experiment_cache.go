package experiment

import (
	"github.com/google/uuid"
)

type activeExperimentCache struct {
	experiments map[uuid.UUID]*activeExperiment
}

func (cache *activeExperimentCache) Set(userId uuid.UUID, experiment *activeExperiment) {
	cache.experiments[userId] = experiment
}

func (cache *activeExperimentCache) Get(userId uuid.UUID) (*activeExperiment, error) {
	experiment, ok := cache.experiments[userId]
	if !ok {
		return nil, errExperimentNotFound
	}
	return experiment, nil
}

func (cache *activeExperimentCache) Delete(userId uuid.UUID) {
	delete(cache.experiments, userId)
}
