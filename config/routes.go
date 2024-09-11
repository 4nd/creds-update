package config

import (
	"creds-update/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func GetRoutes(h *handlers.RouteHandler) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.Handle(
		"/assets/*",
		http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))),
	)

	r.Get("/", h.Welcome)

	r.Route("/creds", func(r chi.Router) {
		r.Get("/", h.ListCreds)
	})

	return r
}
