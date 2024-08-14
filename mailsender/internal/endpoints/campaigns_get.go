package endpoints

import (
	"net/http"
)

func (h *Handler) CampaignGet(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	data, err := h.CampaignService.Get()
	return data, http.StatusOK, err
}
