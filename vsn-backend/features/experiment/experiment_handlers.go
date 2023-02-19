package experiment

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/seg491X-team36/vsn-backend/domain/model"
	"github.com/seg491X-team36/vsn-backend/features/security"
)

// ExperimentService interface required by ExperimentHandlers
type ExperimentService interface {
	Pending(userId uuid.UUID) (experimentId uuid.UUID, pending int)
	StartExperiment(userId, experimentId uuid.UUID) (*model.ExperimentConfig, error)
	StartRound(userId uuid.UUID) (*ExperimentStatus, error)
	StopRound(userId uuid.UUID) (*ExperimentStatus, error)
	Record(userId uuid.UUID, request recordDataRequest) error
}

type ExperimentHandlers struct {
	ExperimentService
}

func (e *ExperimentHandlers) Pending() http.HandlerFunc {
	return postRequestWrapper(func(ctx context.Context, _ pendingExperimentsRequest) pendingExperimentsResponse {
		token, _ := security.AuthToken(ctx)
		experimentId, pending := e.ExperimentService.Pending(token.UserId) // experiment service pending method

		return pendingExperimentsResponse{
			ExperimentId: &experimentId, // could make this nil when pending is 0...
			Pending:      pending,
		}
	})
}

func (e *ExperimentHandlers) StartExperiment() http.HandlerFunc {
	return postRequestWrapper(func(ctx context.Context, req startExperimentRequest) startExperimentResponse {
		token, _ := security.AuthToken(ctx)
		config, err := e.ExperimentService.StartExperiment(token.UserId, req.ExperimentId) // experiment service start experiment method

		return startExperimentResponse{
			Experiment: config,
			Error:      errWrapper(err),
		}
	})
}

func (e *ExperimentHandlers) StartRound() http.HandlerFunc {
	return postRequestWrapper(func(ctx context.Context, req startRoundRequest) startRoundResponse {
		token, _ := security.AuthToken(ctx)
		status, err := e.ExperimentService.StartRound(token.UserId) // experiment service start round method

		return startRoundResponse{
			Status: status,
			Error:  errWrapper(err),
		}
	})
}

func (e *ExperimentHandlers) StopRound() http.HandlerFunc {
	return postRequestWrapper(func(ctx context.Context, req startRoundRequest) stopRoundResponse {
		token, _ := security.AuthToken(ctx)
		status, err := e.ExperimentService.StopRound(token.UserId) // experiment service stop round method

		return stopRoundResponse{
			Status: status,
			Error:  errWrapper(err),
		}
	})
}

func (e *ExperimentHandlers) Record() http.HandlerFunc {
	return postRequestWrapper(func(ctx context.Context, req recordDataRequest) recordDataResponse {
		token, _ := security.AuthToken(ctx)
		err := e.ExperimentService.Record(token.UserId, req) // experiment service record method

		return recordDataResponse{
			Error: errWrapper(err),
		}
	})
}
