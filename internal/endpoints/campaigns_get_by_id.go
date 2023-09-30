package endpoints

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *Handler) CampaignsGetByID(_ http.ResponseWriter, r *http.Request) (responseData interface{}, status int, err error) {
	id := chi.URLParam(r, "id")
	responseData, err = h.CampaignService.GetByID(id)

	return responseData, http.StatusOK, err
}
