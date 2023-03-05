package server

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/seg491X-team36/vsn-backend/codegen/graph"
	"github.com/seg491X-team36/vsn-backend/features/experiment"
	"github.com/seg491X-team36/vsn-backend/features/security"
)

type Server struct {
	UserVerifier         security.UserVerifier
	ExperimentHandlers   *experiment.ExperimentHandlers
	VerificationHandlers *experiment.VerificationHandlers
	Resolvers            graph.ResolverRoot
}

func (s *Server) Handler() http.Handler {
	r := chi.NewRouter()

	graphqlHandler := handler.NewDefaultServer(graph.NewExecutableSchema(
		graph.Config{Resolvers: s.Resolvers},
	))

	// graphql endpoint
	r.Post("/graphql", func(w http.ResponseWriter, r *http.Request) {
		graphqlHandler.ServeHTTP(w, r)
	})

	// graphql playground endpoint
	r.Get("/graphql", playground.Handler("vsn-playground", "/graphql"))

	// experiment api
	r.Route("/experiment", func(r chi.Router) {
		// verification routes
		r.Post("/verify-email", s.VerificationHandlers.EnterEmail())
		r.Post("/verify-code", s.VerificationHandlers.EnterVerificationCode())

		// experiment routes
		r.Group(func(r chi.Router) {
			r.Use(security.Middleware(s.UserVerifier))
			// routes
			r.Post("/start", s.ExperimentHandlers.StartExperiment())
			r.Post("/pending", s.ExperimentHandlers.Pending())
			r.Post("/round/start", s.ExperimentHandlers.StartRound())
			r.Post("/round/stop", s.ExperimentHandlers.StopRound())
			r.Post("/record", s.ExperimentHandlers.Record())
		})
	})

	return r
}
