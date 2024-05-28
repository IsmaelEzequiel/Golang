package endpoints

import (
	"emailSender/internal/domain/campaign"
	internalerrors "emailSender/internal/internalErrors"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Handler_Errors_When_Endpoint_Returns_Erros(t *testing.T) {
	assert := assert.New(t)

	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, http.StatusInternalServerError, internalerrors.ErrInternal
	}

	handlerFunc := HandlerError(endpoint)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusInternalServerError, res.Code)
	assert.Contains(res.Body.String(), internalerrors.ErrInternal.Error())
}

func Test_Handler_Errors_When_Endpoint_Returns_Bad_Request(t *testing.T) {
	assert := assert.New(t)

	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, http.StatusBadRequest, errors.New("Some error")
	}

	handlerFunc := HandlerError(endpoint)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusBadRequest, res.Code)
	assert.Contains(res.Body.String(), "Some error")
}

func Test_Handler_Errors_When_Endpoint_Returns_Object(t *testing.T) {
	assert := assert.New(t)

	objForTest := campaign.Campaign{Name: "nome do ismael"}

	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return objForTest, http.StatusOK, nil
	}

	handlerFunc := HandlerError(endpoint)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	assert.Contains(res.Body.String(), "nome do ismael")
}
