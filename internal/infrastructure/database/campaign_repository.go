package database

import (
	"emailSender/internal/domain/campaign"

	"gorm.io/gorm"
)

type CampaignRepository struct {
	Database *gorm.DB
}

func (c *CampaignRepository) Save(campaign *campaign.Campaign) error {
	xt := c.Database.Save(&campaign)
	return xt.Error
}

func (c *CampaignRepository) Get() ([]campaign.Campaign, error) {
	var campaigns []campaign.Campaign
	xt := c.Database.Find(&campaigns)
	return campaigns, xt.Error
}

func (c *CampaignRepository) GetBy(id string) (*campaign.Campaign, error) {
	var campaign campaign.Campaign
	xt := c.Database.Find(&campaign, "id = ?", id)
	return &campaign, xt.Error
}
