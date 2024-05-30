package endpoints

import (
	"bytes"
	"context"
	"emailSender/internal/contract"
	internalMock "emailSender/internal/test/internal-mock"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setup(body contract.NewCampaign, createdBy string) (*http.Request, *httptest.ResponseRecorder) {
	var buffer bytes.Buffer

	json.NewEncoder((&buffer)).Encode(body)
	req, _ := http.NewRequest("POST", "/", &buffer)
	ctx := context.WithValue(req.Context(), emailKey, createdBy)
	req = req.WithContext(ctx)
	rr := httptest.NewRecorder()

	return req, rr
}

func (h *Handler) Test_CampaignPost_should_save_new_campaign(t *testing.T) {
	assert := assert.New(t)

	body := contract.NewCampaign{
		Name:      "test",
		Content:   "content",
		CreatedBy: "email@email",
		Emails:    []string{"teste@teste.com"},
	}
	service := new(internalMock.CampaignServiceMock)

	service.On("Create", mock.MatchedBy(func(request contract.NewCampaign) bool {
		return request.Name == body.Name && request.CreatedBy == "email@email.com"
	})).Return("123", nil)
	handler := Handler{CampaignService: service}

	req, rr := setup(body, "email@email.com")

	_, status, err := handler.CampaignPost(rr, req)

	assert.Equal(http.StatusCreated, status)
	assert.Nil(err)
}

func (h *Handler) Test_CampaignPost_should_inform_error(t *testing.T) {
	assert := assert.New(t)

	body := contract.NewCampaign{
		Name:      "test",
		Content:   "content",
		CreatedBy: "email@email",
		Emails:    []string{"teste@teste.com"},
	}
	service := new(internalMock.CampaignServiceMock)
	handler := Handler{CampaignService: service}

	service.On("Create", mock.Anything).Return("", errors.New("some error"))

	req, rr := setup(body, "email@email.com")

	json, status, err := handler.CampaignPost(rr, req)

	assert.Equal(http.StatusInternalServerError, status)
	assert.NotNil(err)
	assert.Nil(json)
}
