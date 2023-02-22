package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/seg491X-team36/vsn-backend/codegen/db"
	"github.com/seg491X-team36/vsn-backend/domain/model"
)

func convertExperimentResult(record db.ExperimentResult) model.ExperimentResult {
	return model.ExperimentResult{
		Id:           record.ID,
		CreatedAt:    record.CreatedAt,
		UserId:       record.UserID,
		ExperimentId: record.ExperimentID,
	}
}

func convertExperimentResults(records []db.ExperimentResult) []model.ExperimentResult {
	experimentResults := make([]model.ExperimentResult, len(records))
	for i, record := range records {
		experimentResults[i] = convertExperimentResult(record)
	}
	return experimentResults
}

type ExperimentResultRepository struct {
	Query *db.Queries
}

func (repository *ExperimentResultRepository) CreateExperimentResult(
	ctx context.Context,
	input model.ExperimentResultInput,
) (model.ExperimentResult, error) {
	experimentResult, err := repository.Query.CreateExperimentResult(ctx, db.CreateExperimentResultParams{
		ID:           input.Id,
		UserID:       input.UserId,
		ExperimentID: input.ExperimentId,
	})
	return convertExperimentResult(experimentResult), err
}

func (repository *ExperimentResultRepository) GetExperimentResultsByExperimentId(
	ctx context.Context,
	experimentId uuid.UUID,
) []model.ExperimentResult {
	experimentResults, _ := repository.Query.GetExperimentResultsByExperimentId(ctx, experimentId)
	return convertExperimentResults(experimentResults)
}

func (repository *ExperimentResultRepository) GetExperimentResultsByUserId(
	ctx context.Context,
	userId uuid.UUID,
) []model.ExperimentResult {
	experimentResults, _ := repository.Query.GetExperimentResultsByUserId(ctx, userId)
	return convertExperimentResults(experimentResults)
}
