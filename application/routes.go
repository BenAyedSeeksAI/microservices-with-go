package application

import (
	"net/http"

	"github.com/BenAyedSeeksAI/micro-serv-go/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	router.Route("/orders", loadOrderRoutes)
	return router
}
func loadOrderRoutes(router chi.Router) {
	orderHandler := &handler.Order{}
	router.Post("/create", orderHandler.Create)
	router.Get("/list", orderHandler.List)
	router.Get("/get/{id}", orderHandler.GetByID)
	router.Put("/update/{id}", orderHandler.UpdateByID)
	router.Delete("/delete/{id}", orderHandler.DeleteByID)

}