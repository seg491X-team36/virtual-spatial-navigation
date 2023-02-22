package postgres

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/seg491X-team36/vsn-backend/codegen/db"
	"github.com/seg491X-team36/vsn-backend/domain/model"
)

func convertExperiment(record db.Experiment) model.Experiment {
	var expConfig model.ExperimentConfig = model.ExperimentConfig{}
	json.Unmarshal(record.Config, &expConfig)
	return model.Experiment{
		Id:          record.ID,
		Name:        record.Name,
		Description: record.Description,
		Config:      expConfig,
	}
}

func convertExperiments(records []db.Experiment) []model.Experiment {
	experiments := make([]model.Experiment, len(records))
	for i, record := range records {
		experiments[i] = convertExperiment(record)
	}
	return experiments
}

type ExperimentRepository struct {
	Query *db.Queries
}

func (repository ExperimentRepository) CreateExperiment(
	ctx context.Context,
	input model.ExperimentInput,
) (model.Experiment, error) {
	experiment, err := repository.Query.CreateExperiment(ctx, db.CreateExperimentParams{
		Name:        input.Name,
		Description: input.Description,
	})
	return convertExperiment(experiment), err
}

func (repository ExperimentRepository) GetExperiment(
	ctx context.Context,
	experimentId uuid.UUID,
) (model.Experiment, error) {
	experiment, err := repository.Query.GetExperiment(ctx, experimentId)
	return convertExperiment(experiment), err
}

func (repository ExperimentRepository) GetExperiments(
	ctx context.Context,
) []model.Experiment {
	experiments, _ := repository.Query.GetExperiments(ctx)
	return convertExperiments(experiments)
}
