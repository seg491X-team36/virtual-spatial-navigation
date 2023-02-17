package experiment

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/seg491X-team36/vsn-backend/features/security"
)

type Endpoints struct {
}

func (e *Endpoints) Bind(r chi.Router) {
	r.Route("/experiment", func(r chi.Router) {
		r.Use(security.AuthMiddleware(nil)) // TODO
		r.Post("/pending", func(w http.ResponseWriter, r *http.Request) {})
		r.Post("/start", func(w http.ResponseWriter, r *http.Request) {})
		r.Post("/round/start", func(w http.ResponseWriter, r *http.Request) {})
		r.Post("/round/end", func(w http.ResponseWriter, r *http.Request) {})
		r.Post("/frame", func(w http.ResponseWriter, r *http.Request) {})
	})
}
