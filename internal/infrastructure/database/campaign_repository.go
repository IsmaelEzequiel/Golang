package database

import (
	"emailSender/internal/domain/campaign"

	"gorm.io/gorm"
)

type CampaignRepository struct {
	Database *gorm.DB
}

func (c *CampaignRepository) Create(campaign *campaign.Campaign) error {
	xt := c.Database.Create(&campaign)
	return xt.Error
}

func (c *CampaignRepository) Update(campaign *campaign.Campaign) error {
	xt := c.Database.Save(&campaign)
	return xt.Error
}

func (c *CampaignRepository) Get() ([]campaign.Campaign, error) {
	var campaigns []campaign.Campaign
	xt := c.Database.Preload("Contacts").Find(&campaigns)
	return campaigns, xt.Error
}

func (c *CampaignRepository) GetBy(id string) (*campaign.Campaign, error) {
	var campaign campaign.Campaign
	xt := c.Database.Preload("Contacts").Where("id = ?", id).First(&campaign)
	return &campaign, xt.Error
}

func (c *CampaignRepository) Delete(campaign *campaign.Campaign) error {
	xt := c.Database.Select("Contacts").Delete(&campaign)
	return xt.Error
}
