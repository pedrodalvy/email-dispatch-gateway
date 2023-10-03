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

	db := database.NewDB()
	campaignRepository := database.NewCampaignRepository(db)
	campaignService := campaign.NewService(campaignRepository)
	handler := endpoints.NewHandler(campaignService)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ok"))
	})

	r.Route("/campaigns", func(r chi.Router) {
		r.Use(endpoints.Auth)

		r.Post("/", endpoints.HandlerError(handler.CampaignsPost))
		r.Get("/{id}", endpoints.HandlerError(handler.CampaignsGetByID))
		r.Patch("/{id}/cancel", endpoints.HandlerError(handler.CampaignsPatchCancelByID))
		r.Delete("/{id}", endpoints.HandlerError(handler.CampaignsDeleteByID))
	})

	http.ListenAndServe(":3000", r)
}
