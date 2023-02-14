package experiment

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/seg491X-team36/vsn-backend/features/security"
)

type experimentHandler[Request, Response any] func(ctx context.Context, request Request) Response

func postRequestWrapper[Request, Response any](
	handler experimentHandler[Request, Response],
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// read the request body
		data, err := io.ReadAll(r.Body)
		_ = r.Body.Close()

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// unmarshall the request
		var request Request
		if err := json.Unmarshal(data, request); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// handle the request
		response := handler(r.Context(), request)

		// marshall the response
		data, err = json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// sucessful response
		_, _ = w.Write(data)
		w.WriteHeader(http.StatusOK)
	}
}

type Endpoints struct {
}

func (e *Endpoints) Bind(r chi.Router) {
	r.Post("/verification/email", postRequestWrapper(func(ctx context.Context, request *verificationEmailRequest) *verificationEmailResponse {
		return &verificationEmailResponse{}
	}))

	r.Post("/verification/code", postRequestWrapper(func(ctx context.Context, request *verificationCodeRequest) *verificationCodeResponse {
		return &verificationCodeResponse{}
	}))

	r.Route("/experiment", func(r chi.Router) {
		r.Use(security.AuthMiddleware(nil)) // TODO
		r.Post("/pending", func(w http.ResponseWriter, r *http.Request) {})
		r.Post("/start", func(w http.ResponseWriter, r *http.Request) {})
		r.Post("/round/start", func(w http.ResponseWriter, r *http.Request) {})
		r.Post("/round/end", func(w http.ResponseWriter, r *http.Request) {})
		r.Post("/frame", func(w http.ResponseWriter, r *http.Request) {})
	})
}
