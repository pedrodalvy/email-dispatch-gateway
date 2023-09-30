package main

import (
	"email-dispatch-gateway/internal/domain/campaign"
	"email-dispatch-gateway/internal/endpoints"
	"email-dispatch-gateway/internal/infrastructure/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	campaignService := campaign.NewService(&database.CampaignRepository{})
	handler := endpoints.NewHandler(campaignService)

	r.Post("/campaigns", endpoints.HandlerError(handler.CampaignsPost))
	r.Get("/campaigns/{id}", endpoints.HandlerError(handler.CampaignsGetByID))

	http.ListenAndServe(":3000", r)
}
