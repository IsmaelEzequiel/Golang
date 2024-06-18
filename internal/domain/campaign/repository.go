package campaign

type Repository interface {
	Update(campaign *Campaign) error
	Create(campaign *Campaign) error
	Get() ([]Campaign, error)
	GetBy(id string) (*Campaign, error)
	Delete(*Campaign) error
	GetCampaignsToBeSent() ([]Campaign, error)
}
