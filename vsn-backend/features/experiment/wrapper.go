package experiment

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type experimentHandler[Request, Response any] func(ctx context.Context, req Request) Response

func errWrapper(err error) *string {
	if err != nil {
		msg := err.Error()
		return &msg
	}
	return nil
}

func postRequestWrapper[Request, Response any](
	handler experimentHandler[Request, Response],
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// read the request body
		data, err := io.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// unmarshall the request
		var request Request
		if err := json.Unmarshal(data, &request); err != nil {
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
