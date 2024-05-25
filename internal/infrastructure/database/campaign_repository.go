package database

import (
	"emailSender/internal/domain/campaign"
	"errors"
)

type CampaignRepository struct {
	campaigns []campaign.Campaign
}

func (c *CampaignRepository) Save(campaign *campaign.Campaign) error {
	c.campaigns = append(c.campaigns, *campaign)

	return nil
}

func (c *CampaignRepository) Get() ([]campaign.Campaign, error) {
	return c.campaigns, nil
}

func (c *CampaignRepository) GetById(id string) (*campaign.Campaign, error) {
	for _, camp := range c.campaigns {
		if camp.ID == id {
			return &camp, nil
		}
	}

	return nil, errors.New("campaign not found")
}
