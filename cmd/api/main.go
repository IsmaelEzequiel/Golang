package main

import (
	"emailSender/internal/domain/campaign"
	"emailSender/internal/endpoints"
	"emailSender/internal/infrastructure/database"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	db := database.NewDB()

	// Routes
	campaignService := campaign.ServiceImpl{Repository: &database.CampaignRepository{Database: db}}
	handler := endpoints.Handler{CampaignService: &campaignService}

	r.Route("/campaigns", func(r chi.Router) {
		r.Post("/", endpoints.HandlerError(handler.CampaignPost))
		r.Get("/", endpoints.HandlerError(handler.CampaignGet))
		r.Get("/{id}", endpoints.HandlerError(handler.CampaignGetBy))
		r.Patch("/cancel/{id}", endpoints.HandlerError(handler.CampaignsUpdate))
	})

	http.ListenAndServe(":3000", r)
}
