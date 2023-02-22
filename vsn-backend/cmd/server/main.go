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
		Resolvers: &resolvers.Root{
			ExperimentResolver:       &resolvers.ExperimentResolvers{},
			ExperimentResultResolver: &resolvers.ExperimentResultResolvers{},
			ExperimentConfigResolver: &resolvers.ExperimentConfigResolvers{},
			InviteResolver:           &resolvers.InviteResolvers{},
			MutationResolver:         &resolvers.MutationResolvers{},
			QueryResolver:            &resolvers.QueryResolvers{},
			UserResolver:             &resolvers.UserResolvers{},
		},
	}

	_ = http.ListenAndServe("localhost:3001", srv.Handler())
}
