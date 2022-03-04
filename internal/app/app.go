package app

import (
	"net/http"
	"refactoring/internal/controller/handler"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	//r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	return r
}

func Route(router *chi.Mux, handle *handler.Handler) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(time.Now().String()))
	})

	router.Route("/api", func(r chi.Router) {

		r.Route("/v1", func(r chi.Router) {

			r.Route("/users", func(r chi.Router) {

				r.Get("/", handle.SearchUsers)
				r.Post("/", handle.CreateUser)

				r.Route("/{id}", func(r chi.Router) {

					r.Get("/", handle.GetUser)
					r.Patch("/", handle.UpdateUser)
					r.Delete("/", handle.DeleteUser)
				})
			})
		})
	})
}
