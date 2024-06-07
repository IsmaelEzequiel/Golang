package endpoints

import (
	"bytes"
	"context"
	internalMock "emailSender/internal/test/internal-mock"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/go-chi/chi/v5"
)

var (
	service *internalMock.CampaignServiceMock
	handler = Handler{}
)

func setup() {
	service = new(internalMock.CampaignServiceMock)
	handler = Handler{CampaignService: service}
}

func newHttpTest(method string, url string, body interface{}) (*http.Request, *httptest.ResponseRecorder) {
	var buffer bytes.Buffer

	if body != nil {
		json.NewEncoder((&buffer)).Encode(body)
	}

	req, _ := http.NewRequest(method, url, &buffer)
	rr := httptest.NewRecorder()

	return req, rr
}

func addParameters(req *http.Request, keyParameter string, valueParameter string) *http.Request {
	chiContext := chi.NewRouteContext()
	chiContext.URLParams.Add(keyParameter, valueParameter)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiContext))

	return req
}
