package endpoints

import (
	"emailSender/internal/contract"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {

	email := r.Context().Value("email").(string)

	var request contract.NewCampaign
	render.DecodeJSON(r.Body, &request)

	fmt.Println(request)

	request.CreatedBy = email
	id, err := h.CampaignService.Create(request)
	return map[string]string{"id": id}, http.StatusCreated, err
}
