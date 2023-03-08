package experiment

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostWrapper(t *testing.T) {
	type dummyRequest struct {
		A int `json:"a"`
		B int `json:"b"`
	}

	type dummyResponse struct {
		C int `json:"c"`
	}

	handler := postRequestWrapper(func(ctx context.Context, req dummyRequest) dummyResponse {
		return dummyResponse{
			C: req.A + req.B,
		}
	})

	t.Run("happy-path", func(t *testing.T) {
		reqData, _ := json.Marshal(dummyRequest{A: 1, B: 2})
		expectedResData, _ := json.Marshal(dummyResponse{C: 3})

		r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqData))
		w := httptest.NewRecorder()

		handler.ServeHTTP(w, r)

		actualResData, _ := io.ReadAll(w.Result().Body)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, expectedResData, actualResData)
	})
}
