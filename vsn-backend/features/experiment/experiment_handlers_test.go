package experiment

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

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
	token string,
	handler http.Handler,
) {
	reqData, _ := json.Marshal(req)

	// send the http request
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqData))
	r.Header.Add("token", token)
	handler.ServeHTTP(w, r)
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
		RoundsTotal: 1,
		Resume:      model.CONTINUE_ROUND,
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

		HandlerTC(t, req, res, token, pending)
	})

	t.Run("start", func(t *testing.T) {
		start := middleware(experimentHandlers.StartExperiment())

		req := startExperimentRequest{
			ExperimentId: experimentId,
		}

		res := startExperimentResponse{
			Data: &startExperimentData{
				Config: experiment.Config,
				Status: model.ExperimentStatus{
					RoundInProgress: false,
					RoundsCompleted: 0,
					RoundsTotal:     experiment.Config.RoundsTotal,
				},
				Frame: nil,
			},
			Error: nil,
		}

		HandlerTC(t, req, res, token, start)
	})

	t.Run("start-round", func(t *testing.T) {
		startRound := middleware(experimentHandlers.StartRound())

		req := startRoundRequest{}

		res := startRoundResponse{
			Status: &model.ExperimentStatus{
				RoundInProgress: true,
				RoundsCompleted: 0,
				RoundsTotal:     experiment.Config.RoundsTotal,
			},
			Error: nil,
		}

		HandlerTC(t, req, res, token, startRound)
	})

	t.Run("record", func(t *testing.T) {
		record := middleware(experimentHandlers.Record())

		req := recordDataRequest{
			Data: experimentData{
				Frames: make([]frame, 5),
				Events: []event{{Name: eventRewardFound, Timestamp: time.Now().UTC()}},
			},
		}

		res := recordDataResponse{
			Error: nil,
		}

		HandlerTC(t, req, res, token, record)
	})

	t.Run("stop-round", func(t *testing.T) {
		stopRound := middleware(experimentHandlers.StopRound())

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

		HandlerTC(t, req, res, token, stopRound)
	})

	// the events were recorded
	assertEventsRecordedOrdered(t, recorder, eventRoundStart, eventRewardFound, eventRoundStop)

	// 5 from the record test case, 5 from the stop round test case
	assert.Equal(t, 10, len(recorder.frames))

}
