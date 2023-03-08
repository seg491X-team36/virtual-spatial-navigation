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
	StartExperiment(ctx context.Context, userId, experimentId uuid.UUID) (*startExperimentData, error)
	StartRound(ctx context.Context, userId uuid.UUID) (*model.ExperimentStatus, error)
	StopRound(ctx context.Context, userId uuid.UUID, data experimentData) (*model.ExperimentStatus, error)
	Record(ctx context.Context, userId uuid.UUID, data experimentData) error
}

type ExperimentHandlers struct {
	ExperimentService
}

func (e *ExperimentHandlers) Pending() http.HandlerFunc {
	return postRequestWrapper(func(ctx context.Context, _ pendingExperimentsRequest) pendingExperimentsResponse {
		claims, _ := security.GetUserClaims(ctx)
		return e.ExperimentService.Pending(ctx, claims.UserId) // experiment service pending method
	})
}

func (e *ExperimentHandlers) StartExperiment() http.HandlerFunc {
	return postRequestWrapper(func(ctx context.Context, req startExperimentRequest) startExperimentResponse {
		claims, _ := security.GetUserClaims(ctx)
		res, err := e.ExperimentService.StartExperiment(ctx, claims.UserId, req.ExperimentId)

		return startExperimentResponse{
			Data:  res,
			Error: errWrapper(err),
		}
	})
}

func (e *ExperimentHandlers) StartRound() http.HandlerFunc {
	return postRequestWrapper(func(ctx context.Context, req startRoundRequest) startRoundResponse {
		claims, _ := security.GetUserClaims(ctx)
		status, err := e.ExperimentService.StartRound(ctx, claims.UserId) // experiment service start round method

		return startRoundResponse{
			Status: status,
			Error:  errWrapper(err),
		}
	})
}

func (e *ExperimentHandlers) StopRound() http.HandlerFunc {
	return postRequestWrapper(func(ctx context.Context, req stopRoundRequest) stopRoundResponse {
		claims, _ := security.GetUserClaims(ctx)
		status, err := e.ExperimentService.StopRound(ctx, claims.UserId, req.Data) // experiment service stop round method

		return stopRoundResponse{
			Status: status,
			Error:  errWrapper(err),
		}
	})
}

func (e *ExperimentHandlers) Record() http.HandlerFunc {
	return postRequestWrapper(func(ctx context.Context, req recordDataRequest) recordDataResponse {
		claims, _ := security.GetUserClaims(ctx)
		err := e.ExperimentService.Record(ctx, claims.UserId, req.Data) // experiment service record method

		return recordDataResponse{
			Error: errWrapper(err),
		}
	})
}
