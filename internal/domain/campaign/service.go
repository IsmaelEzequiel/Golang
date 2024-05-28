package campaign

import (
	"emailSender/internal/contract"
	internalerrors "emailSender/internal/internalErrors"
)

type Service interface {
	Get() ([]Campaign, error)
	GetById(id string) (*contract.CampaignResponse, error)
	Create(newCampaign contract.NewCampaign) (string, error)
}

type ServiceImpl struct {
	Repository Repository
}

func (s *ServiceImpl) Create(newCampaign contract.NewCampaign) (string, error) {
	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

	if err != nil {
		return "", err
	}

	err = s.Repository.Save(campaign)

	if err != nil {
		return "", internalerrors.ErrInternal
	}

	return campaign.ID, err
}

func (s *ServiceImpl) Get() ([]Campaign, error) {
	campaign, err := s.Repository.Get()

	if err != nil {
		return nil, internalerrors.ErrInternal
	}

	return campaign, nil
}

func (s *ServiceImpl) GetById(id string) (*contract.CampaignResponse, error) {
	campaign, err := s.Repository.GetById(id)

	if err != nil {
		return nil, err
	}

	return &contract.CampaignResponse{
		ID:      campaign.ID,
		Name:    campaign.Name,
		Content: campaign.Content,
		Status:  campaign.Status,
	}, nil
}
