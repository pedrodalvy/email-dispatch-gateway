package campaign

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var (
	name    = "Campaign Name"
	content = "Campaign Content"
	emails  = []string{"a@domain.com", "b@domain.com"}
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	// ACT
	campaign, err := NewCampaign(name, content, emails)

	// ASSERT
	require.Equal(t, name, campaign.Name)
	require.Equal(t, content, campaign.Content)
	require.Equal(t, len(emails), len(campaign.Contacts))
	require.Empty(t, err)
}

func Test_NewCampaign_IDIsNotEmpty(t *testing.T) {
	// ACT
	campaign, _ := NewCampaign(name, content, emails)

	// ASSERT
	require.NotEmpty(t, campaign.ID)
}

func Test_NewCampaign_CreatedOnMustBeNow(t *testing.T) {
	// ARRANGE
	now := time.Now().Add(-time.Minute)

	// ACT
	campaign, _ := NewCampaign(name, content, emails)

	// ASSERT
	require.Greater(t, campaign.CreatedOn, now)
}

func Test_NewCampaign_MustValidateName(t *testing.T) {
	// ACT
	campaign, err := NewCampaign("", content, emails)

	// ASSERT
	require.Equal(t, "name is required", err.Error())
	require.Empty(t, campaign)
}

func Test_NewCampaign_MustValidateContent(t *testing.T) {
	// ACT
	campaign, err := NewCampaign(name, "", emails)

	// ASSERT
	require.Equal(t, "content is required", err.Error())
	require.Empty(t, campaign)
}

func Test_NewCampaign_MustValidateContacts(t *testing.T) {
	// ACT
	campaign, err := NewCampaign(name, content, []string{})

	// ASSERT
	require.Equal(t, "contacts is required", err.Error())
	require.Empty(t, campaign)
}
