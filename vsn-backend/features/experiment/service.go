package experiment

import (
	"context"

	"github.com/google/uuid"
	"github.com/seg491X-team36/vsn-backend/domain/model"
	"github.com/seg491X-team36/vsn-backend/features/experiment/experimentsync"
)

type inviteRepository interface {
	GetPendingInvites(ctx context.Context, userId uuid.UUID) []model.Invite
}

type experimentRepository interface {
	GetExperiment(ctx context.Context, experimentId uuid.UUID) (model.Experiment, error)
}

type experimentResultRepository interface {
	CreateExperimentResult(ctx context.Context, input model.ExperimentResultInput) (model.ExperimentResult, error)
}

type Service struct {
	invites           inviteRepository
	experiments       experimentRepository
	experimentResults experimentResultRepository
	activeExperiments *activeExperimentCache
	recorderFactory   recorderFactory
}

func NewService(
	invites inviteRepository,
	experiments experimentRepository,
	experimentResults experimentResultRepository,
	factory recorderFactory,
) ExperimentService {
	return &syncExperimentService{
		UserMap: experimentsync.NewMap(),
		service: &Service{
			invites:           invites,
			experiments:       experiments,
			experimentResults: experimentResults,
			activeExperiments: &activeExperimentCache{
				experiments: map[uuid.UUID]*activeExperiment{},
			},
			recorderFactory: factory,
		},
	}
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

func (s *Service) StartExperiment(ctx context.Context, userId, experimentId uuid.UUID) (*startExperimentData, error) {
	// get the pending experiment
	res := s.Pending(ctx, userId)

	// pending experiment in progress
	if res.ExperimentInProgress {
		return s.ResumeExperiment(ctx, userId, experimentId)
	}

	// no pending experiment
	if res.ExperimentId == nil {
		return nil, errExperimentNotFound
	}

	// pending experiment id and experiment id do not match
	if *res.ExperimentId != experimentId {
		return nil, errExperimentNotFound
	}

	experiment, _ := s.experiments.GetExperiment(ctx, *res.ExperimentId)

	// create the active experiment
	activeExperiment := &activeExperiment{
		TrackingId:   uuid.New(), // assign a new tracking id
		ExperimentId: experimentId,
		UserId:       userId,

		ExperimentStatus: model.NewExperimentStatus(experiment.Config.RoundsTotal),
		ExperimentConfig: experiment.Config,

		// create the recorder
		recorder: s.recorderFactory(recorderParams{
			ExperimentId: experimentId,
			TrackingId:   experimentId,
		}),
	}

	s.activeExperiments.Set(userId, activeExperiment)

	return &startExperimentData{
		Experiment: activeExperiment.ExperimentConfig,
		Status:     activeExperiment.ExperimentStatus,
		Frame:      nil,
	}, nil
}

func (s *Service) ResumeExperiment(ctx context.Context, userId, experimentId uuid.UUID) (*startExperimentData, error) {
	experiment, _ := s.activeExperiments.Get(userId)
	frame := experiment.Resume()

	return &startExperimentData{
		Experiment: experiment.ExperimentConfig,
		Status:     experiment.ExperimentStatus,
		Frame:      frame,
	}, nil
}

func (s *Service) StartRound(ctx context.Context, userId uuid.UUID) (*model.ExperimentStatus, error) {
	// get the active experiment
	experiment, err := s.activeExperiments.Get(userId)
	if err != nil {
		return nil, err
	}

	// call start round
	status, err := experiment.StartRound()
	return &status, err
}

func (s *Service) StopRound(ctx context.Context, userId uuid.UUID, data experimentData) (*model.ExperimentStatus, error) {
	// get the active experiment
	experiment, err := s.activeExperiments.Get(userId)
	if err != nil {
		return nil, err
	}

	// call stop round
	status, err := experiment.StopRound(data)

	// the experiment is done
	if status.Done() {
		s.activeExperiments.Delete(userId)
		_, _ = s.experimentResults.CreateExperimentResult(ctx, model.ExperimentResultInput{
			Id:           experiment.TrackingId,
			UserId:       experiment.UserId,
			ExperimentId: experiment.ExperimentId,
		})
	}

	return &status, err
}

func (s *Service) Record(ctx context.Context, userId uuid.UUID, data experimentData) error {
	// get the active experiment
	experiment, err := s.activeExperiments.Get(userId)
	if err != nil {
		return err
	}

	experiment.Record(data)
	return nil
}
