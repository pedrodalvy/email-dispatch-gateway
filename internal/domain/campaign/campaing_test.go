package campaign

import (
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var (
	name      = "Campaign Name"
	content   = "Campaign Content"
	emails    = []string{"a@domain.com", "b@domain.com"}
	createdBy = "test@email.com"
	fake      = faker.New()
)

func Test_Campaign_NewCampaign(t *testing.T) {
	t.Run("should create campaign", func(t *testing.T) {
		// ACT
		campaign, err := NewCampaign(name, content, emails, createdBy)

		// ASSERT
		require.Equal(t, name, campaign.Name)
		require.Equal(t, content, campaign.Content)
		require.Equal(t, len(emails), len(campaign.Contacts))
		require.Equal(t, createdBy, campaign.CreatedBy)
		require.Empty(t, err)
	})

	t.Run("should return a campaign ID", func(t *testing.T) {
		// ACT
		campaign, _ := NewCampaign(name, content, emails, createdBy)

		// ASSERT
		require.NotEmpty(t, campaign.ID)
	})

	t.Run("should return a valid createdOn time", func(t *testing.T) {
		// ARRANGE
		now := time.Now().Add(-time.Minute)

		// ACT
		campaign, _ := NewCampaign(name, content, emails, createdBy)

		// ASSERT
		require.Greater(t, campaign.CreatedOn, now)
	})

	t.Run("should validate if name has less than 5 characters", func(t *testing.T) {
		// ACT
		campaign, err := NewCampaign("four", content, emails, createdBy)

		// ASSERT
		require.Equal(t, "name must be more than 5 characters", err.Error())
		require.Empty(t, campaign)
	})

	t.Run("should validate if name has more than 24 characters", func(t *testing.T) {
		// ACT
		invalidName := fake.Lorem().Text(25)
		campaign, err := NewCampaign(invalidName, content, emails, createdBy)

		// ASSERT
		require.Equal(t, "name must be less than 24 characters", err.Error())
		require.Empty(t, campaign)
	})

	t.Run("should validate if content has less than 5 characters", func(t *testing.T) {
		// ACT
		campaign, err := NewCampaign(name, "four", emails, createdBy)

		// ASSERT
		require.Equal(t, "content must be more than 5 characters", err.Error())
		require.Empty(t, campaign)
	})

	t.Run("should validate if content has more than 1024 characters", func(t *testing.T) {
		// ACT
		invalidContent := fake.Lorem().Text(1040)
		campaign, err := NewCampaign(name, invalidContent, emails, createdBy)

		// ASSERT
		require.Equal(t, "content must be less than 1024 characters", err.Error())
		require.Empty(t, campaign)
	})

	t.Run("should validate if contacts has less than 1", func(t *testing.T) {
		// ACT
		campaign, err := NewCampaign(name, content, []string{}, createdBy)

		// ASSERT
		require.Equal(t, "contacts must be greater than or equal to 1", err.Error())
		require.Empty(t, campaign)
	})

	t.Run("should validate if contacts has an invalid email", func(t *testing.T) {
		// ACT
		campaign, err := NewCampaign(name, content, []string{"invalid email"}, createdBy)

		// ASSERT
		require.Equal(t, "email is invalid", err.Error())
		require.Empty(t, campaign)
	})

	t.Run("should create a campaign with status pending", func(t *testing.T) {
		// ACT
		campaign, _ := NewCampaign(name, content, emails, createdBy)

		// ASSERT
		require.Equal(t, Pending, campaign.Status)
	})

	t.Run("should validate if createdBy is an invalid email", func(t *testing.T) {
		// ACT
		campaign, err := NewCampaign(name, content, emails, "invalid_field")

		// ASSERT
		require.Equal(t, "createdby is invalid", err.Error())
		require.Empty(t, campaign)
	})
}

func Test_Campaign_Cancel(t *testing.T) {
	t.Run("should cancel a campaign", func(t *testing.T) {
		// ARRANGE
		campaign, _ := NewCampaign(name, content, emails, createdBy)

		// ACT
		err := campaign.Cancel()

		// ASSERT
		require.Nil(t, err)
		require.Equal(t, Canceled, campaign.Status)
	})

	t.Run("should return an error if campaign status is invalid", func(t *testing.T) {
		// ARRANGE
		campaign, _ := NewCampaign(name, content, emails, createdBy)
		campaign.Status = "another"

		// ACT
		err := campaign.Cancel()

		// ASSERT
		require.Equal(t, "campaign status is invalid", err.Error())
	})
}

func Test_Campaign_Delete(t *testing.T) {
	t.Run("should delete a campaign", func(t *testing.T) {
		// ARRANGE
		campaign, _ := NewCampaign(name, content, emails, createdBy)

		// ACT
		err := campaign.Delete()

		// ASSERT
		require.Nil(t, err)
		require.Equal(t, Deleted, campaign.Status)
	})

	t.Run("should delete a cancelled campaign", func(t *testing.T) {
		// ARRANGE
		campaign, _ := NewCampaign(name, content, emails, createdBy)
		campaign.Cancel()

		// ACT
		err := campaign.Delete()

		// ASSERT
		require.Nil(t, err)
		require.Equal(t, Deleted, campaign.Status)
	})

	t.Run("should return an error if campaign status is invalid", func(t *testing.T) {
		// ARRANGE
		campaign, _ := NewCampaign(name, content, emails, createdBy)
		campaign.Status = "another"

		// ACT
		err := campaign.Delete()

		// ASSERT
		require.Equal(t, "campaign status is invalid", err.Error())
	})
}
