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

func (r *repositoryMock) Get() ([]Campaign, error) {
	return nil, nil
}

func (r *repositoryMock) GetById(id string) (*Campaign, error) {
	args := r.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Campaign), nil
}

var (
	newCampaign = contract.NewCampaign{
		Name:    "Ismael",
		Content: "Content",
		Emails:  []string{"a@a.com", "b@b.com"},
	}
	service = ServiceImpl{}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(nil)

	service := ServiceImpl{Repository: repositoryMock}

	id, err := service.Create(newCampaign)

	assert.NotEmpty(id)
	assert.Nil(err)
}

func Test_Fail_To_Create_Campaign(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(errors.New(""))

	service := ServiceImpl{Repository: repositoryMock}

	id, err := service.Create(newCampaign)

	assert.Empty(id)
	assert.Equal("Internal Server Error", err.Error())
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

	service := ServiceImpl{Repository: repositoryMock}

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

	service := ServiceImpl{Repository: repositoryMock}

	_, err := service.Create(newCampaign)

	assert.NotNil(err)
	assert.True(errors.Is(internalerrors.ErrInternal, err))
}

func Test_Create_Error_SaveDomain(t *testing.T) {
	assert := assert.New(t)

	_, err := service.Create(contract.NewCampaign{})

	assert.NotNil(err)
	assert.False(errors.Is(internalerrors.ErrInternal, err))
}

func Test_GetById_campaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

	repositoryMock := new(repositoryMock)
	repositoryMock.On("GetById", mock.MatchedBy(func(id string) bool {
		return id == campaign.ID
	})).Return(campaign, nil)
	service.Repository = repositoryMock

	returnedCampaign, _ := service.GetById(campaign.ID)

	assert.Equal(campaign.ID, returnedCampaign.ID)
	assert.Equal(campaign.Name, returnedCampaign.Name)
	assert.Equal(campaign.Content, returnedCampaign.Content)
}

func Test_GetById_Error_campaign(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(repositoryMock)
	repositoryMock.On("GetById", mock.Anything).Return(nil, errors.New("campaign not found"))
	service.Repository = repositoryMock

	_, err := service.GetById("123")

	assert.NotNil(err)
	assert.Equal(err.Error(), "campaign not found")
}
