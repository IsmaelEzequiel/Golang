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

	// Routes
	service := campaign.Service{Repository: &database.CampaignRepository{}}
	handler := endpoints.Handler{CampaignService: service}

	r.Post("/campaings", endpoints.HandlerError(handler.CampaignPost))
	r.Get("/campaings", endpoints.HandlerError(handler.CampaignGet))

	http.ListenAndServe(":3000", r)
}
