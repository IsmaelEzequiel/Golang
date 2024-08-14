package endpoints

import (
	internalMock "emailSender/internal/test/internal-mock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCampaign_delete_success(t *testing.T) {
	service := new(internalMock.CampaignServiceMock)
	handler := Handler{CampaignService: service}
	service.On("Delete", mock.Anything).Return(nil)

	req, _ := http.NewRequest("PATCH", "/", nil)
	rr := httptest.NewRecorder()

	_, status, err := handler.CampaignsDelete(rr, req)

	assert.Nil(t, err)
	assert.Equal(t, status, http.StatusOK)
}
