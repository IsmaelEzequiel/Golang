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
	assert := assert.New(t)

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

	_, status, _ := handler.CampaignGetBy(rr, req)

	assert.Equal(status, http.StatusOK)
	// assert.Equal(body.ID, response.(*contract.CampaignResponse).ID)
	// assert.Equal(body.Name, response.(*contract.CampaignResponse).Name)
	// assert.Equal(body.Content, response.(*contract.CampaignResponse).Content)
	// assert.Equal(body.Status, response.(*contract.CampaignResponse).Status)
}

func Test_CampaignGet_should_return_error_when_something_went_wrong(t *testing.T) {
	assert := assert.New(t)

	service := new(internalMock.CampaignServiceMock)
	service.On("GetBy", mock.Anything).Return(nil, errors.New("Some error"))
	handler := Handler{CampaignService: service}
	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	_, status, err := handler.CampaignGetBy(rr, req)

	assert.Equal(status, http.StatusOK)
	assert.Equal(err, "some error")
}
