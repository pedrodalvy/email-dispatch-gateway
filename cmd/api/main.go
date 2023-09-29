package main

import (
	"email-dispatch-gateway/internal/contract"
	"email-dispatch-gateway/internal/domain/campaign"
	"email-dispatch-gateway/internal/infrastructure/database"
	internalerrors "email-dispatch-gateway/internal/internal-errors"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	repository := &database.CampaignRepository{}
	service := campaign.NewService(repository)

	r.Post("/campaigns", func(w http.ResponseWriter, r *http.Request) {
		var campaignDTO contract.NewCampaignDTO
		render.DecodeJSON(r.Body, &campaignDTO)

		id, err := service.Create(campaignDTO)
		if err != nil {
			status := http.StatusBadRequest
			if errors.Is(err, internalerrors.ErrInternalServerError) {
				status = http.StatusInternalServerError
			}

			render.Status(r, status)
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}

		render.Status(r, http.StatusCreated)
		render.JSON(w, r, map[string]string{"id": id})
	})

	http.ListenAndServe(":3000", r)
}
