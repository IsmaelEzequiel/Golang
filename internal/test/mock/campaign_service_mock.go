package mock

import (
	"emailSender/internal/contract"
	"emailSender/internal/domain/campaign"

	"github.com/stretchr/testify/mock"
)

type CampaignServiceMock struct {
	mock.Mock
}

func (r *CampaignServiceMock) Create(newCampaign contract.NewCampaign) (string, error) {
	args := r.Called(newCampaign)
	return args.String(0), args.Error(0)
}

func (r *CampaignServiceMock) GetById(id string) (*contract.CampaignResponse, error) {
	// args := r.Called(newCampaign)
	return nil, nil
}

func (r *CampaignServiceMock) Get() ([]campaign.Campaign, error) {
	// args := r.Called(newCampaign)
	return nil, nil
}
