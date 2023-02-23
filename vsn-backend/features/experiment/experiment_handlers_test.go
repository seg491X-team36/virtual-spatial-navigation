package experiment

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/seg491X-team36/vsn-backend/domain/model"
	"github.com/seg491X-team36/vsn-backend/features/security"
	"github.com/seg491X-team36/vsn-backend/features/security/verification"
	"github.com/stretchr/testify/assert"
)

func HandlerTC[Request, Response any](
	t *testing.T,
	req Request,
	resExpected Response,
	target string,
	expectedStatus int,
	token string,
	handler http.HandlerFunc,
) {
	reqData, _ := json.Marshal(req)

	// send the http request
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, target, bytes.NewBuffer(reqData))
	r.Header.Add("token", token)
	handler(w, r)
	res := w.Result()
	defer res.Body.Close()

	// read the response
	resData, _ := io.ReadAll(res.Body)
	var resActual Response
	_ = json.Unmarshal(resData, &resActual)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, resExpected, resActual)
}

func TestExperimentHandlers(t *testing.T) {
	userId := uuid.New()
	experimentId := uuid.New()

	recorder := &recorderStub{}

	experiment := model.Experiment{Id: experimentId, Config: model.ExperimentConfig{
		RoundsTotal:  2,
		ResumeConfig: model.CONTINUE_ROUND,
	}}

	service := NewService(
		&inviteRepositoryStub{
			Invites: []model.Invite{{ID: uuid.New(), UserID: userId, ExperimentID: experimentId}},
		},
		&experimentRepositoryStub{
			Experiment: experiment,
		},
		&experimentResultRepositoryStub{},
		recorderStubFactory(recorder),
	)

	experimentHandlers := &ExperimentHandlers{
		ExperimentService: service,
	}

	tokens := &verification.TokenManager{
		Secret: []byte("secret"),
	}

	token := tokens.Generate(verification.Token{
		UserId: userId,
	})

	middleware := security.AuthMiddleware(tokens)

	t.Run("pending", func(t *testing.T) {
		pending := middleware(experimentHandlers.Pending())

		req := pendingExperimentsRequest{}
		res := pendingExperimentsResponse{
			ExperimentId:         &experimentId,
			ExperimentInProgress: false,
			Pending:              1,
		}
		HandlerTC(t, req, res, "/pending", http.StatusOK, token, pending.ServeHTTP)
	})

	t.Run("start", func(t *testing.T) {
		start := middleware(experimentHandlers.StartExperiment())

		req := startExperimentRequest{
			ExperimentId: experimentId,
		}

		res := startExperimentResponse{
			Data: &startExperimentData{
				Experiment: experiment.Config,
				Status: model.ExperimentStatus{
					RoundInProgress: false,
					RoundsCompleted: 0,
					RoundsTotal:     experiment.Config.RoundsTotal,
				},
				Frame: nil,
			},
			Error: nil,
		}

		HandlerTC(t, req, res, "/start", http.StatusOK, token, start.ServeHTTP)
	})

	t.Run("round-start", func(t *testing.T) {
		roundStart := middleware(experimentHandlers.StartRound())

		req := startRoundRequest{}

		res := startRoundResponse{
			Status: &model.ExperimentStatus{
				RoundInProgress: true,
				RoundsCompleted: 0,
				RoundsTotal:     experiment.Config.RoundsTotal,
			},
			Error: nil,
		}

		HandlerTC(t, req, res, "/round/start", http.StatusOK, token, roundStart.ServeHTTP)
	})

	t.Run("round-stop", func(t *testing.T) {
		roundStop := middleware(experimentHandlers.StopRound())

		req := stopRoundRequest{
			Data: experimentData{
				Frames: make([]frame, 5), // 5 empty frames
			},
		}

		res := stopRoundResponse{
			Status: &model.ExperimentStatus{
				RoundInProgress: false,
				RoundsCompleted: 1,
				RoundsTotal:     experiment.Config.RoundsTotal,
			},
			Error: nil,
		}

		HandlerTC(t, req, res, "/round/stop", http.StatusOK, token, roundStop.ServeHTTP)
	})

	t.Run("record", func(t *testing.T) {
		record := middleware(experimentHandlers.Record())

		req := recordDataRequest{
			Data: experimentData{
				Frames: make([]frame, 5),
				Events: make([]event, 1),
			},
		}

		res := recordDataResponse{
			Error: nil,
		}

		HandlerTC(t, req, res, "/record", http.StatusOK, token, record.ServeHTTP)
	})

	// the record test case includes 1 event, at minimum 1 start round, and 1 stop round should also be recorded
	assert.GreaterOrEqual(t, len(recorder.events), 3)

	// the record and stop round test case record 5 frames each
	assert.Equal(t, len(recorder.frames), 10)
}
