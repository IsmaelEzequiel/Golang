package endpoints

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignStart_200(t *testing.T) {
	setup()
	service.On("Start", mock.MatchedBy(func(id string) bool {
		return id == "123"
	})).Return(nil)

	req, rr := newHttpTest("PATCH", "/", nil)
	req = addParameters(req, "id", "123")

	_, status, err := handler.CampaignStart(rr, req)

	assert.Equal(t, status, http.StatusOK)
	assert.Nil(t, err)
}

func Test_CampaignStart_Err(t *testing.T) {
	setup()

	service.On("Start", mock.Anything).Return(errors.New("Some error"))
	req, rr := newHttpTest("PATCH", "/", nil)

	_, _, err := handler.CampaignStart(rr, req)

	assert.Equal(t, "Some error", err.Error())
}
