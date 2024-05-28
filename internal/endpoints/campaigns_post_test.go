package endpoints

import (
	"bytes"
	"emailSender/internal/contract"
	"emailSender/internal/domain/campaign"
	internalMock "emailSender/internal/test/mock"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	body = contract.NewCampaign{
		Name:    "test",
		Content: "content",
		Emails:  []string{"teste@teste.com"},
	}
	service = new(internalMock.CampaignServiceMock)
	handler = Handler{CampaignService: &campaign.ServiceImpl{}}
	buffer  bytes.Buffer
)

func (h *Handler) Test_CampaignPost_should_save_new_campaign(t *testing.T) {
	assert := assert.New(t)

	service.On("Create", mock.MatchedBy(func(request contract.NewCampaign) bool {
		return request.Name == body.Name
	})).Return("123", nil)

	json.NewEncoder((&buffer)).Encode(body)
	req, _ := http.NewRequest("POST", "/", &buffer)
	rr := httptest.NewRecorder()

	_, status, err := handler.CampaignPost(rr, req)

	assert.Equal(http.StatusCreated, status)
	assert.Nil(err)
}

func (h *Handler) Test_CampaignPost_should_inform_error(t *testing.T) {
	assert := assert.New(t)

	service.On("Create", mock.Anything).Return("", errors.New("some error"))

	json.NewEncoder((&buffer)).Encode(body)
	req, _ := http.NewRequest("POST", "/", &buffer)
	rr := httptest.NewRecorder()

	json, status, err := handler.CampaignPost(rr, req)

	assert.Equal(http.StatusInternalServerError, status)
	assert.NotNil(err)
	assert.Nil(json)
}
