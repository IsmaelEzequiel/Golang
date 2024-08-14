package main

import (
	"emailSender/internal/domain/campaign"
	"emailSender/internal/endpoints"
	"emailSender/internal/infrastructure/database"
	"emailSender/internal/infrastructure/mail"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	r := chi.NewRouter()

	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading .env file")
	}

	// Middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	db := database.NewDB()

	// Routes
	campaignService := campaign.ServiceImpl{
		Repository: &database.CampaignRepository{Database: db},
		SendEmail:  mail.SendEmail,
	}
	handler := endpoints.Handler{CampaignService: &campaignService}

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Route("/campaigns", func(r chi.Router) {
		r.Use(endpoints.CheckAuthMiddleware)
		r.Post("/", endpoints.HandlerError(handler.CampaignPost))
		r.Get("/", endpoints.HandlerError(handler.CampaignGet))
		r.Get("/{id}", endpoints.HandlerError(handler.CampaignGetBy))
		r.Patch("/delete/{id}", endpoints.HandlerError(handler.CampaignsDelete))
		r.Patch("/start/{id}", endpoints.HandlerError(handler.CampaignStart))
	})

	http.ListenAndServe(":3000", r)
}
