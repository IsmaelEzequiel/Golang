package endpoints

import (
	"emailSender/internal/contract"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setup2(body contract.NewCampaign, createdBy string) (*http.Request, *httptest.ResponseRecorder) {
	req, rr := newHttpTest("POST", "/", body)
	req = addParameters(req, "email", createdBy)

	return req, rr
}

func (h *Handler) Test_CampaignPost_should_save_new_campaign(t *testing.T) {
	setup()
	body := contract.NewCampaign{
		Name:      "test",
		Content:   "content",
		CreatedBy: "email@email",
		Emails:    []string{"teste@teste.com"},
	}

	service.On("Create", mock.MatchedBy(func(request contract.NewCampaign) bool {
		return request.Name == body.Name && request.CreatedBy == "email@email.com"
	})).Return("123", nil)

	req, rr := setup2(body, "email@email.com")

	_, status, err := handler.CampaignPost(rr, req)

	assert.Equal(t, http.StatusCreated, status)
	assert.Nil(t, err)
}

func (h *Handler) Test_CampaignPost_should_inform_error(t *testing.T) {
	setup()
	body := contract.NewCampaign{
		Name:      "test",
		Content:   "content",
		CreatedBy: "email@email",
		Emails:    []string{"teste@teste.com"},
	}

	service.On("Create", mock.Anything).Return("", errors.New("some error"))

	req, rr := setup2(body, "email@email.com")

	json, status, err := handler.CampaignPost(rr, req)

	assert.Equal(t, http.StatusInternalServerError, status)
	assert.NotNil(t, err)
	assert.Nil(t, json)
}
