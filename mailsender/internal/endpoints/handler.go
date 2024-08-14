package endpoints

import "emailSender/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}
