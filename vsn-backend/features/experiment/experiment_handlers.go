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
	Pending(ctx context.Context, userId uuid.UUID) pendingExperimentsResponse
	StartExperiment(ctx context.Context, userId, experimentId uuid.UUID) startExperimentResponse
	StartRound(ctx context.Context, userId uuid.UUID) (*model.ExperimentStatus, error)
	StopRound(ctx context.Context, userId uuid.UUID) (*model.ExperimentStatus, error)
	Record(ctx context.Context, userId uuid.UUID, request recordDataRequest) error
}

type ExperimentHandlers struct {
	ExperimentService
}

func (e *ExperimentHandlers) Pending() http.HandlerFunc {
	return postRequestWrapper(func(ctx context.Context, _ pendingExperimentsRequest) pendingExperimentsResponse {
		token, _ := security.AuthToken(ctx)
		return e.ExperimentService.Pending(ctx, token.UserId) // experiment service pending method
	})
}

func (e *ExperimentHandlers) StartExperiment() http.HandlerFunc {
	return postRequestWrapper(func(ctx context.Context, req startExperimentRequest) startExperimentResponse {
		token, _ := security.AuthToken(ctx)
		return e.ExperimentService.StartExperiment(ctx, token.UserId, req.ExperimentId) // experiment service start experiment method
	})
}

func (e *ExperimentHandlers) StartRound() http.HandlerFunc {
	return postRequestWrapper(func(ctx context.Context, req startRoundRequest) startRoundResponse {
		token, _ := security.AuthToken(ctx)
		status, err := e.ExperimentService.StartRound(ctx, token.UserId) // experiment service start round method

		return startRoundResponse{
			Status: status,
			Error:  errWrapper(err),
		}
	})
}

func (e *ExperimentHandlers) StopRound() http.HandlerFunc {
	return postRequestWrapper(func(ctx context.Context, req startRoundRequest) stopRoundResponse {
		token, _ := security.AuthToken(ctx)
		status, err := e.ExperimentService.StopRound(ctx, token.UserId) // experiment service stop round method

		return stopRoundResponse{
			Status: status,
			Error:  errWrapper(err),
		}
	})
}

func (e *ExperimentHandlers) Record() http.HandlerFunc {
	return postRequestWrapper(func(ctx context.Context, req recordDataRequest) recordDataResponse {
		token, _ := security.AuthToken(ctx)
		err := e.ExperimentService.Record(ctx, token.UserId, req) // experiment service record method

		return recordDataResponse{
			Error: errWrapper(err),
		}
	})
}
