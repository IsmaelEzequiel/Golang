package campaign

import (
	internalerrors "emailSender/internal/internalErrors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	ID         string `gorm:"size:50" json:"id"`
	Email      string `validate:"email" json:"email"`
	CampaignID string `json:"campaign_id"`
}

const (
	Pending  = "pending"
	Started  = "started"
	Canceled = "canceled"
	Done     = "done"
)

type Campaign struct {
	ID        string    `validate:"required" gorm:"size:50" json:"id"`
	Name      string    `validate:"min=5,max=24" gorm:"size:100" json:"name"`
	CreatedAt time.Time `validate:"required" json:"created_at"`
	Content   string    `validate:"min=5,max=1024" json:"content"`
	Contacts  []Contact `validate:"min=1,dive" json:"contacts"`
	Status    string    `json:"status"`
}

func (c *Campaign) Cancel() {
	c.Status = Canceled
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	contacts := make([]Contact, len(emails))

	for index, email := range emails {
		contacts[index].Email = email
		contacts[index].ID = xid.New().String()
	}

	newCampaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedAt: time.Now(),
		Content:   content,
		Contacts:  contacts,
		Status:    Pending,
	}

	err := internalerrors.ValidateStruct(newCampaign)

	if err != nil {
		return nil, err
	}

	return newCampaign, nil
}
