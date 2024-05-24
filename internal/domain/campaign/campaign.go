package campaign

import (
	internalerrors "emailSender/internal/internalErrors"
	"time"

	"github.com/rs/xid"
)

type Contacts struct {
	Email string `validate:"email"`
}

type Campaign struct {
	ID        string     `validate:"required"`
	Name      string     `validate:"min=5,max=24"`
	CreatedAt time.Time  `validate:"required"`
	Content   string     `validate:"min=5,max=1024"`
	Contacts  []Contacts `validate:"min=1,dive"`
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	contacts := make([]Contacts, len(emails))

	for index, email := range emails {
		contacts[index].Email = email
	}

	newCampaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedAt: time.Now(),
		Content:   content,
		Contacts:  contacts,
	}

	err := internalerrors.ValidateStruct(newCampaign)

	if err != nil {
		return nil, err
	}

	return newCampaign, nil
}
