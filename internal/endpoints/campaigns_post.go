package endpoints

import (
	"email-dispatch-gateway/internal/contract"
	"github.com/go-chi/render"
	"net/http"
)

func (h *Handler) CampaignsPost(_ http.ResponseWriter, r *http.Request) (responseData interface{}, status int, err error) {
	var campaignDTO contract.NewCampaignDTO
	render.DecodeJSON(r.Body, &campaignDTO)

	id, err := h.CampaignService.Create(campaignDTO)

	return map[string]string{"id": id}, http.StatusCreated, err
}
