package main

import (
	"net/http"

	"github.com/seg491X-team36/vsn-backend/features/experiment"
	"github.com/seg491X-team36/vsn-backend/features/resolvers"
	"github.com/seg491X-team36/vsn-backend/server"
	"github.com/spf13/afero"
)

func main() {
	recorder := experiment.NewFSRecorderFactory(
		afero.NewBasePathFs(afero.NewOsFs(), "./data"),
	)

	srv := &server.Server{
		ExperimentHandlers: &experiment.ExperimentHandlers{
			ExperimentService: experiment.NewService(nil, nil, recorder), // TODO
		},
		VerificationHandlers: &experiment.VerificationHandlers{
			// TODO
		},
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
