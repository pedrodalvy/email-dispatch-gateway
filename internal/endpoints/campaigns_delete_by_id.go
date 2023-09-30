package endpoints

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *Handler) CampaignsDeleteByID(_ http.ResponseWriter, r *http.Request) (responseData interface{}, status int, err error) {
	id := chi.URLParam(r, "id")
	err = h.CampaignService.DeleteByID(id)

	return nil, http.StatusNoContent, err
}
