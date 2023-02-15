package server

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/seg491X-team36/vsn-backend/codegen/graph"
	"github.com/seg491X-team36/vsn-backend/features/experiment"
)

type Server struct {
	Experiments experiment.Endpoints
	Resolvers   graph.ResolverRoot
}

func (s *Server) Handler() http.Handler {
	router := chi.NewRouter()

	graphqlHandler := handler.NewDefaultServer(graph.NewExecutableSchema(
		graph.Config{Resolvers: s.Resolvers},
	))

	router.Post("/graphql", func(w http.ResponseWriter, r *http.Request) {
		graphqlHandler.ServeHTTP(w, r)
	})

	router.Get("/graphql", playground.Handler("vsn-playground", "/graphql"))

	s.Experiments.Bind(router)

	return router
}
