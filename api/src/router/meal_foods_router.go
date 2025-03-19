package router

import (
	"github.com/go-chi/chi/v5"

	"toyfit/src/handlers"
)

func SetMealFoodsRouter() *chi.Mux {
	r := chi.NewRouter()

	var h handlers.MealFoodsHandler
	r.Post("/", h.Create)
	// r.Get("/{id:[0-9]+}", jobPost.GetById)
	// r.Get("/", jobPost.GetAll)
	// r.Put("/{id:[0-9]+}/address", jobPost.SetAddress)
	// r.Put("/{id:[0-9]+}", jobPost.Update)
	return r
}
