package experiment

import (
	"context"

	"github.com/google/uuid"
	"github.com/seg491X-team36/vsn-backend/domain/model"
	"github.com/seg491X-team36/vsn-backend/features/experiment/experimentsync"
)

type syncExperimentService struct {
	*experimentsync.Map
	service ExperimentService
}

func (s *syncExperimentService) Pending(ctx context.Context, userId uuid.UUID) pendingExperimentsResponse {
	lock := s.Lock(userId)
	defer lock.Unlock()

	return s.service.Pending(ctx, userId)
}

func (s *syncExperimentService) StartExperiment(ctx context.Context, userId, experimentId uuid.UUID) (*startExperimentData, error) {
	lock := s.Lock(userId)
	defer lock.Unlock()

	return s.service.StartExperiment(ctx, userId, experimentId)
}

func (s *syncExperimentService) StartRound(ctx context.Context, userId uuid.UUID) (*model.ExperimentStatus, error) {
	lock := s.Lock(userId)
	defer lock.Unlock()

	return s.service.StartRound(ctx, userId)
}

func (s *syncExperimentService) StopRound(ctx context.Context, userId uuid.UUID, data experimentData) (*model.ExperimentStatus, error) {
	lock := s.Lock(userId)
	defer lock.Unlock()

	return s.service.StopRound(ctx, userId, data)
}

func (s *syncExperimentService) Record(ctx context.Context, userId uuid.UUID, data experimentData) error {
	lock := s.Lock(userId)
	defer lock.Unlock()

	return s.service.Record(ctx, userId, data)
}
