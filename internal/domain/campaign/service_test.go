package campaign

import (
	"emailSender/internal/contract"
	internalerrors "emailSender/internal/internalErrors"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

var (
	newCampaign = contract.NewCampaign{
		Name:    "Ismael",
		Content: "Content",
		Emails:  []string{"a@a.com", "b@b.com"},
	}
	service = Service{}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(nil)

	service := Service{Repository: repositoryMock}

	id, err := service.Create(newCampaign)

	assert.NotEmpty(id)
	assert.Nil(err)
}

func Test_Fail_To_Create_Campaign(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(errors.New(""))

	service := Service{Repository: repositoryMock}

	id, err := service.Create(newCampaign)

	assert.Empty(id)
	assert.Equal("internal server error", err.Error())
}

func Test_Create_SaveCampaign(t *testing.T) {
	newCampaign := contract.NewCampaign{
		Name:    "Ismael",
		Content: "Content",
		Emails:  []string{"a@a.com", "b@b.com"},
	}

	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name || campaign.Content != newCampaign.Content || len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}

		return true
	})).Return(nil)

	service := Service{Repository: repositoryMock}

	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)
}

func Test_Create_SaveCampaignErrorDB(t *testing.T) {
	assert := assert.New(t)

	newCampaign := contract.NewCampaign{
		Name:    "Ismael",
		Content: "Content",
		Emails:  []string{"a@a.com", "b@b.com"},
	}

	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(internalerrors.ErrInternal)

	service := Service{Repository: repositoryMock}

	_, err := service.Create(newCampaign)

	assert.NotNil(err)
	assert.True(errors.Is(internalerrors.ErrInternal, err))
}

func Test_Create_Error_SaveDomain(t *testing.T) {
	assert := assert.New(t)

	newCampaign.Name = ""

	_, err := service.Create(newCampaign)

	assert.NotNil(err)
	assert.Equal("name is lower then 5", err.Error())
}
