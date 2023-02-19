package main

import (
	"net/http"

	"github.com/seg491X-team36/vsn-backend/features/experiment"
	"github.com/seg491X-team36/vsn-backend/features/resolvers"
	"github.com/seg491X-team36/vsn-backend/server"
)

func main() {
	srv := &server.Server{
		ExperimentHandlers:   &experiment.ExperimentHandlers{},
		VerificationHandlers: &experiment.VerificationHandlers{},
		Resolvers:            &resolvers.Root{},
	}

	_ = http.ListenAndServe("localhost:3001", srv.Handler())
}
