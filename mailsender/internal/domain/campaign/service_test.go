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
	repositoryMock  *internalmock.CampaignRepositoryMock
	service         = campaign.ServiceImpl{}
	campaignPending *campaign.Campaign
)

func setup() {
	repositoryMock = new(internalmock.CampaignRepositoryMock)
	service.Repository = repositoryMock
	campaignPending = &campaign.Campaign{ID: "1", Status: campaign.Pending}
}

func SetupGetBy(campaign *campaign.Campaign) {
	repositoryMock.On("GetBy", mock.Anything).Return(campaign, nil)
}

func Test_Create_Campaign(t *testing.T) {
	setup()
	assert := assert.New(t)

	repositoryMock.On("Create", mock.Anything).Return(nil)

	id, err := service.Create(newCampaign)

	assert.NotEmpty(id)
	assert.Nil(err)
}

func Test_Fail_To_Create_Campaign(t *testing.T) {
	setup()
	assert := assert.New(t)

	repositoryMock.On("Create", mock.Anything).Return(errors.New("Error"))

	id, err := service.Create(newCampaign)

	assert.Empty(id)
	assert.Equal("Internal Server Error", err.Error())
}

func Test_Create_SaveCampaign(t *testing.T) {
	setup()
	newCampaign := contract.NewCampaign{
		Name:      "Ismael",
		Content:   "Content",
		CreatedBy: "email@email.com",
		Emails:    []string{"a@a.com", "b@b.com"},
	}

	repositoryMock.On("Create", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		if campaign.Name != newCampaign.Name || campaign.Content != newCampaign.Content || len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}

		return true
	})).Return(nil)

	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)
}

func Test_Create_SaveCampaignErrorDB(t *testing.T) {
	setup()
	assert := assert.New(t)

	newCampaign := contract.NewCampaign{
		Name:      "Ismael",
		Content:   "Content",
		CreatedBy: "email@email.com",
		Emails:    []string{"a@a.com", "b@b.com"},
	}

	repositoryMock.On("Create", mock.Anything).Return(internalerrors.ErrInternal)

	_, err := service.Create(newCampaign)

	assert.NotNil(err)
	assert.True(errors.Is(internalerrors.ErrInternal, err))
}

func Test_Create_Error_SaveDomain(t *testing.T) {
	setup()
	assert := assert.New(t)

	_, err := service.Create(contract.NewCampaign{})

	assert.NotNil(err)
	assert.False(errors.Is(internalerrors.ErrInternal, err))
}

func Test_GetBy_campaign(t *testing.T) {
	setup()
	assert := assert.New(t)

	campaign, _ := campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)

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
	setup()
	assert := assert.New(t)

	repositoryMock.On("GetBy", mock.Anything).Return(nil, errors.New(""))
	service.Repository = repositoryMock

	_, err := service.GetBy("123")

	assert.NotNil(err)
	assert.Equal(err.Error(), "Internal Server Error")
}

func Test_Delete_return_not_found(t *testing.T) {
	setup()
	assert := assert.New(t)

	repositoryMock.On("GetBy", mock.Anything).Return(nil, gorm.ErrRecordNotFound)
	service.Repository = repositoryMock

	err := service.Delete("invalid key")

	assert.Equal(err.Error(), gorm.ErrRecordNotFound.Error())
}

func Test_Delete_return_status_invalid(t *testing.T) {
	setup()
	assert := assert.New(t)

	campaign := &campaign.Campaign{ID: "123", Status: campaign.Started}

	repositoryMock.On("GetBy", mock.MatchedBy(func(id string) bool {
		return id == campaign.ID
	})).Return(campaign, nil)
	service.Repository = repositoryMock

	err := service.Delete(campaign.ID)

	assert.Equal(err.Error(), "status invalid to be deleted")
}

func Test_Delete_return_success(t *testing.T) {
	setup()
	assert := assert.New(t)

	repositoryMock.On("GetBy", mock.MatchedBy(func(id string) bool {
		return id == campaignPending.ID
	})).Return(campaignPending, nil)
	repositoryMock.On("Delete", mock.Anything).Return(nil)
	service.Repository = repositoryMock

	err := service.Delete(campaignPending.ID)

	assert.Nil(err)
}

func Test_Delete_return_general_error(t *testing.T) {
	setup()
	assert := assert.New(t)

	campaignFound, _ := campaign.NewCampaign("ismael", "content", []string{"email@email.com"}, newCampaign.CreatedBy)

	SetupGetBy(campaignFound)

	repositoryMock.On("Delete", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		return campaignFound == campaign
	})).Return(errors.New("internal error"))

	service.Repository = repositoryMock

	err := service.Delete(campaignFound.ID)

	assert.NotNil(err)
}

func Test_Delete_returnNil_on_delete(t *testing.T) {
	setup()
	assert := assert.New(t)

	campaignFound, _ := campaign.NewCampaign("ismael", "content", []string{"email@email.com"}, newCampaign.CreatedBy)

	SetupGetBy(campaignFound)

	repositoryMock.On("Delete", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		return campaignFound == campaign
	})).Return(nil)

	service.Repository = repositoryMock

	err := service.Delete(campaignFound.ID)

	assert.Nil(err)
}

func Test_Start_sending_email_invalid_campaign(t *testing.T) {
	setup()
	assert := assert.New(t)

	repositoryMock.On("GetBy", mock.Anything).Return(nil, gorm.ErrRecordNotFound)
	service.Repository = repositoryMock

	err := service.Start("invalid key")

	assert.Equal(err.Error(), gorm.ErrRecordNotFound.Error())
}

func Test_Start_return_InvalidStatus(t *testing.T) {
	setup()
	assert := assert.New(t)

	campaignFound := &campaign.Campaign{ID: "1", Status: campaign.Started}
	repositoryMock.On("GetBy", mock.Anything).Return(campaignFound, nil)
	service.Repository = repositoryMock

	err := service.Start(campaignFound.ID)

	assert.Equal("status invalid to be deleted", err.Error())
}

func Test_Start_should_send_email(t *testing.T) {
	setup()
	SetupGetBy(campaignPending)
	assert := assert.New(t)

	repositoryMock.On("Update", mock.Anything).Return(nil)
	service.Repository = repositoryMock
	sentEmail := false

	sendEmail := func(campaign *campaign.Campaign) error {
		sentEmail = true
		return nil
	}

	service.SendEmail = sendEmail

	service.Start(campaignPending.ID)

	assert.True(sentEmail)
}

func Test_Start_should_error_sending_email(t *testing.T) {
	setup()
	SetupGetBy(campaignPending)
	assert := assert.New(t)

	service.Repository = repositoryMock

	sendEmail := func(campaign *campaign.Campaign) error {
		return errors.New("error sending email")
	}

	service.SendEmail = sendEmail

	err := service.Start(campaignPending.ID)

	assert.Equal(err.Error(), internalerrors.ErrInternal.Error())
}

func Test_Start_should_change_status_to_done(t *testing.T) {
	setup()
	SetupGetBy(campaignPending)
	assert := assert.New(t)

	repositoryMock.On("Update", mock.MatchedBy(func(campaignUpdate *campaign.Campaign) bool {
		return campaignUpdate.ID == campaignPending.ID && campaignUpdate.Status == campaign.Done
	})).Return(nil)
	service.Repository = repositoryMock

	sendEmail := func(campaignParams *campaign.Campaign) error {
		return nil
	}

	service.SendEmail = sendEmail

	service.Start(campaignPending.ID)

	assert.Equal(campaignPending.Status, campaign.Done)
}

func Test_Start_should_change_status_to_error(t *testing.T) {
	setup()
	SetupGetBy(campaignPending)
	assert := assert.New(t)

	repositoryMock.On("Update", mock.Anything).Return(errors.New("some error"))
	service.Repository = repositoryMock

	sendEmail := func(campaignParams *campaign.Campaign) error {
		return nil
	}

	service.SendEmail = sendEmail

	err := service.Start(campaignPending.ID)

	assert.Equal(err, internalerrors.ErrInternal)
}
