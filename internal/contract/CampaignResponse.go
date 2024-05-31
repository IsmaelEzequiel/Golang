package contract

type CampaignResponse struct {
	ID                   string `json:"id"`
	Name                 string `json:"name"`
	Content              string `json:"content"`
	Status               string `json:"status"`
	AmountOfEmailsToSend int    `json:"amount_of_emails_to_send"`
	CreatedBy            string `json:"created_by"`
}
