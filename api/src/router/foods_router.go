package router

import (
	"github.com/go-chi/chi/v5"

	"toyfit/src/handlers"
)

func SetFoodsRouter() *chi.Mux {
	r := chi.NewRouter()

	var h handlers.FoodsHandler
	// r.Post("/", h.Create)
	// r.Get("/{id:[0-9]+}", jobPost.GetById)
	r.Get("/", h.GetAll)
	// r.Put("/{id:[0-9]+}/address", jobPost.SetAddress)
	// r.Put("/{id:[0-9]+}", jobPost.Update)
	return r
}
