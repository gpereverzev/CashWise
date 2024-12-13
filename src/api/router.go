package api

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gpereverzev/CashWise/src/api/handlers"
	"github.com/gpereverzev/CashWise/src/data"
	"net/http"
	"time"
)

type Config struct {
	MasterDB data.MasterDB
}

func Run(config Config) {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), handlers.MasterDBContextKey, config.MasterDB)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api/cashwise", func(r chi.Router) {
		r.Route("/v1/login", func(r chi.Router) {
			r.Post("/registration", handlers.Registration)
			r.Get("/auth", handlers.Auth)
		})

		r.Route("/v1/user", func(r chi.Router) {
			r.Get("/profile", handlers.GetUserProfile)
		})
	})

	http.ListenAndServe(":8000", r)
}
