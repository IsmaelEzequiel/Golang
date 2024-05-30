package campaign_test

import (
	"emailSender/internal/contract"
	"emailSender/internal/domain/campaign"
	internalerrors "emailSender/internal/internalErrors"
	internalmock "emailSender/internal/test/internal-mock"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	newCampaign = contract.NewCampaign{
		Name:      "Ismael",
		Content:   "Content",
		CreatedBy: "ismael@ismael.com",
		Emails:    []string{"a@a.com", "b@b.com"},
	}
	service = campaign.ServiceImpl{}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("Create", mock.Anything).Return(nil)

	service := campaign.ServiceImpl{Repository: repositoryMock}

	id, err := service.Create(newCampaign)

	assert.NotEmpty(id)
	assert.Nil(err)
}

func Test_Fail_To_Create_Campaign(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("Create", mock.Anything).Return(errors.New("Error"))

	service := campaign.ServiceImpl{Repository: repositoryMock}

	id, err := service.Create(newCampaign)

	assert.Empty(id)
	assert.Equal("Internal Server Error", err.Error())
}

func Test_Create_SaveCampaign(t *testing.T) {
	newCampaign := contract.NewCampaign{
		Name:      "Ismael",
		Content:   "Content",
		CreatedBy: "email@email.com",
		Emails:    []string{"a@a.com", "b@b.com"},
	}

	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("Create", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		if campaign.Name != newCampaign.Name || campaign.Content != newCampaign.Content || len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}

		return true
	})).Return(nil)

	service := campaign.ServiceImpl{Repository: repositoryMock}

	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)
}

func Test_Create_SaveCampaignErrorDB(t *testing.T) {
	assert := assert.New(t)

	newCampaign := contract.NewCampaign{
		Name:      "Ismael",
		Content:   "Content",
		CreatedBy: "email@email.com",
		Emails:    []string{"a@a.com", "b@b.com"},
	}

	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("Create", mock.Anything).Return(internalerrors.ErrInternal)

	service := campaign.ServiceImpl{Repository: repositoryMock}

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

func Test_GetBy_campaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)

	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("GetBy", mock.MatchedBy(func(id string) bool {
		return id == campaign.ID
	})).Return(campaign, nil)
	service.Repository = repositoryMock

	returnedCampaign, _ := service.GetBy(campaign.ID)

	assert.Equal(campaign.ID, returnedCampaign.ID)
	assert.Equal(campaign.Name, returnedCampaign.Name)
	assert.Equal(campaign.Content, returnedCampaign.Content)
}

func Test_GetBy_Error_campaign(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("GetBy", mock.Anything).Return(nil, errors.New(""))
	service.Repository = repositoryMock

	_, err := service.GetBy("123")

	assert.NotNil(err)
	assert.Equal(err.Error(), "Internal Server Error")
}

func Test_Delete_return_not_found(t *testing.T) {
	assert := assert.New(t)

	campaignIdInvalid := "invalid key"
	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("GetBy", mock.Anything).Return(nil, gorm.ErrRecordNotFound)
	service.Repository = repositoryMock

	err := service.Delete(campaignIdInvalid)

	assert.Equal(err.Error(), gorm.ErrRecordNotFound.Error())
}

func Test_Delete_return_status_invalid(t *testing.T) {
	assert := assert.New(t)

	campaign := &campaign.Campaign{ID: "123", Status: campaign.Started}

	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("GetBy", mock.MatchedBy(func(id string) bool {
		return id == campaign.ID
	})).Return(campaign, nil)
	service.Repository = repositoryMock

	err := service.Delete(campaign.ID)

	assert.Equal(err.Error(), "status invalid to be deleted")
}

// func Test_Delete_return_success(t *testing.T) {
// 	assert := assert.New(t)

// 	campaign := &campaign.Campaign{ID: "123", Status: campaign.Pending}

// 	repositoryMock := new(internalmock.CampaignRepositoryMock)
// 	repositoryMock.On("GetBy", mock.MatchedBy(func(id string) bool {
// 		return id == campaign.ID
// 	})).Return(campaign, nil)
// 	service.Repository = repositoryMock

// 	err := service.Delete(campaign.ID)

// 	assert.Nil(err)
// }

func Test_Delete_return_general_error(t *testing.T) {
	assert := assert.New(t)

	campaignFound, _ := campaign.NewCampaign("ismael", "content", []string{"email@email.com"}, newCampaign.CreatedBy)

	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("GetBy", mock.Anything).Return(campaignFound, nil)
	repositoryMock.On("Delete", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		return campaignFound == campaign
	})).Return(errors.New("internal error"))

	service.Repository = repositoryMock

	err := service.Delete(campaignFound.ID)

	assert.NotNil(err)
}

func Test_Delete_returnNil_on_delete(t *testing.T) {
	assert := assert.New(t)

	campaignFound, _ := campaign.NewCampaign("ismael", "content", []string{"email@email.com"}, newCampaign.CreatedBy)

	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("GetBy", mock.Anything).Return(campaignFound, nil)
	repositoryMock.On("Delete", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		return campaignFound == campaign
	})).Return(nil)

	service.Repository = repositoryMock

	err := service.Delete(campaignFound.ID)

	assert.Nil(err)
}
