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

func Test_Campaign_NewCampaign(t *testing.T) {
	t.Run("should create campaign", func(t *testing.T) {
		// ACT
		campaign, err := NewCampaign(name, content, emails)

		// ASSERT
		require.Equal(t, name, campaign.Name)
		require.Equal(t, content, campaign.Content)
		require.Equal(t, len(emails), len(campaign.Contacts))
		require.Empty(t, err)
	})

	t.Run("should return a campaign ID", func(t *testing.T) {
		// ACT
		campaign, _ := NewCampaign(name, content, emails)

		// ASSERT
		require.NotEmpty(t, campaign.ID)
	})

	t.Run("should return a valid createdOn time", func(t *testing.T) {
		// ARRANGE
		now := time.Now().Add(-time.Minute)

		// ACT
		campaign, _ := NewCampaign(name, content, emails)

		// ASSERT
		require.Greater(t, campaign.CreatedOn, now)
	})

	t.Run("should validate name", func(t *testing.T) {
		// ACT
		campaign, err := NewCampaign("", content, emails)

		// ASSERT
		require.Equal(t, "name is required", err.Error())
		require.Empty(t, campaign)
	})

	t.Run("should validate content", func(t *testing.T) {
		// ACT
		campaign, err := NewCampaign(name, "", emails)

		// ASSERT
		require.Equal(t, "content is required", err.Error())
		require.Empty(t, campaign)
	})

	t.Run("should validate contacts", func(t *testing.T) {
		// ACT
		campaign, err := NewCampaign(name, content, []string{})

		// ASSERT
		require.Equal(t, "contacts is required", err.Error())
		require.Empty(t, campaign)
	})
}
