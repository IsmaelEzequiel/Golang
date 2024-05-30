package campaign

import (
	"emailSender/internal/contract"
	internalerrors "emailSender/internal/internalErrors"
	"errors"
)

type Service interface {
	Get() ([]Campaign, error)
	GetBy(id string) (*contract.CampaignResponse, error)
	Delete(id string) error
	Create(newCampaign contract.NewCampaign) (string, error)
}

type ServiceImpl struct {
	Repository Repository
}

func (s *ServiceImpl) Create(newCampaign contract.NewCampaign) (string, error) {
	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)

	if err != nil {
		return "", err
	}

	err = s.Repository.Create(campaign)

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

func (s *ServiceImpl) GetBy(id string) (*contract.CampaignResponse, error) {
	campaign, err := s.Repository.GetBy(id)

	if err != nil {
		return nil, internalerrors.ProcessErrorNotFound(err)
	}

	return &contract.CampaignResponse{
		ID:                   campaign.ID,
		Name:                 campaign.Name,
		Content:              campaign.Content,
		Status:               campaign.Status,
		AmountOfEmailsToSend: len(campaign.Contacts),
	}, nil
}

func (s *ServiceImpl) Delete(id string) error {
	campaign, err := s.Repository.GetBy(id)

	if err != nil {
		return internalerrors.ProcessErrorNotFound(err)
	}

	if campaign.Status != Pending {
		return errors.New("status invalid to be deleted")
	}

	campaign.Delete()
	err = s.Repository.Delete(campaign)

	if err != nil {
		return internalerrors.ErrInternal
	}

	return nil
}
