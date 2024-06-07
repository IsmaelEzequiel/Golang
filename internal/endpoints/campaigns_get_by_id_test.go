package endpoints

import (
	"emailSender/internal/contract"
	internalMock "emailSender/internal/test/internal-mock"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignGet_should_return_campaign(t *testing.T) {
	body := contract.CampaignResponse{
		ID:      "123",
		Name:    "test",
		Content: "content",
		Status:  "pending",
	}

	service := new(internalMock.CampaignServiceMock)
	service.On("GetBy", mock.Anything).Return(&body, nil)
	handler := Handler{CampaignService: service}
	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	response, status, _ := handler.CampaignGetBy(rr, req)

	assert.Equal(t, status, http.StatusOK)
	assert.Equal(t, body.ID, response.(*contract.CampaignResponse).ID)
	assert.Equal(t, body.Name, response.(*contract.CampaignResponse).Name)
	assert.Equal(t, body.Content, response.(*contract.CampaignResponse).Content)
	assert.Equal(t, body.Status, response.(*contract.CampaignResponse).Status)
}

func Test_CampaignGet_should_return_error_when_something_went_wrong(t *testing.T) {
	service := new(internalMock.CampaignServiceMock)
	service.On("GetBy", mock.Anything).Return(nil, errors.New("Some error"))
	handler := Handler{CampaignService: service}
	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	_, status, err := handler.CampaignGetBy(rr, req)

	assert.Equal(t, status, http.StatusOK)
	assert.Equal(t, err.Error(), "Some error")
}
