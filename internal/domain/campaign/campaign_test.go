package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
)

var (
	name     = "campaign x"
	content  = "bodys"
	contacts = []string{"email@email.com", "email2@email.com"}
	now      = time.Now().Add(-time.Minute)
	fake     = faker.New()
)

func Test_NewCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}

func Test_NewCampaign_IDIsNotNil(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreatedAtMustBeNow(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Greater(campaign.CreatedAt, now)
}

func Test_NewCampaign_MustValidateNameMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, contacts)

	assert.Equal("name is lower then 5", err.Error())
}

func Test_NewCampaign_MustValidateNameMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(fake.Lorem().Text(25), content, contacts)

	assert.Equal("name is greater then 24", err.Error())
}

func Test_NewCampaign_MustValidateContentMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", contacts)

	assert.Equal("content is lower then 5", err.Error())
}

func Test_NewCampaign_MustValidateContentMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, fake.Lorem().Text(1040), contacts)

	assert.Equal("content is greater then 1024", err.Error())
}

func Test_NewCampaign_MustValidateContactsMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, nil)

	assert.Equal("contacts is lower then 1", err.Error())
}

func Test_NewCampaign_MustValidateContactsInvalid(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{"email_invalid"})

	assert.Equal("email is invalid", err.Error())
}
