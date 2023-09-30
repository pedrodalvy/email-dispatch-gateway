package endpoints

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *Handler) CampaignsGet(_ http.ResponseWriter, r *http.Request) (campaignResponse interface{}, status int, err error) {
	id := chi.URLParam(r, "id")
	campaignResponse, err = h.CampaignService.GetByID(id)

	return campaignResponse, http.StatusOK, err
}
