package endpoints

import (
	"email-dispatch-gateway/internal/contract"
	"github.com/go-chi/render"
	"net/http"
)

func (h *Handler) CampaignsPost(_ http.ResponseWriter, r *http.Request) (responseData interface{}, status int, err error) {
	var campaignDTO contract.NewCampaignDTO
	render.DecodeJSON(r.Body, &campaignDTO)

	userEmail := r.Context().Value("email").(string)
	campaignDTO.CreatedBy = userEmail

	id, err := h.CampaignService.Create(campaignDTO)
	if err == nil {
		responseData = map[string]string{"id": id}
	}

	return responseData, http.StatusCreated, err
}
