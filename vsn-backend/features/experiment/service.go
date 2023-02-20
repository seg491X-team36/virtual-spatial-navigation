package experiment

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/seg491X-team36/vsn-backend/domain/model"
)

type inviteRepository interface {
	GetPendingInvites(ctx context.Context, userId uuid.UUID) []model.Invite
}

type Service struct {
	invites           inviteRepository
	activeExperiments *activeExperimentCache
}

func (s *Service) Pending(ctx context.Context, userId uuid.UUID) pendingExperimentsResponse {
	experiment, _ := s.activeExperiments.Get(userId)
	invites := s.invites.GetPendingInvites(ctx, userId)

	// there's an experiment in progress
	if experiment != nil {
		return pendingExperimentsResponse{
			ExperimentId:         &experiment.ExperimentId,
			ExperimentInProgress: true,
			Pending:              len(invites) - 1, // the in progress experiment will be in the pending invites
		}
	}

	// there's at least one pending experiment
	if len(invites) > 0 {
		return pendingExperimentsResponse{
			ExperimentId:         &invites[0].ExperimentID,
			ExperimentInProgress: false,
			Pending:              len(invites),
		}
	}

	// there are no pending experiments
	return pendingExperimentsResponse{
		ExperimentId:         nil,
		ExperimentInProgress: false,
		Pending:              len(invites), // 0
	}
}

func (s *Service) StartExperiment(userId, experimentId uuid.UUID) startExperimentResponse {
	return startExperimentResponse{}
}

func (s *Service) StartRound(userId uuid.UUID) (*model.ExperimentStatus, error) {
	return nil, errors.New("not implemented")
}

func (s *Service) StopRound(userId uuid.UUID) (*model.ExperimentStatus, error) {
	return nil, errors.New("not implemented")
}

func (s *Service) Record(userId uuid.UUID, request recordDataRequest) error {
	return errors.New("not implemented")
}
