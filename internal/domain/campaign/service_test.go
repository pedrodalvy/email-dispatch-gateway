package campaign_test

import (
	"email-dispatch-gateway/internal/contract"
	"email-dispatch-gateway/internal/domain/campaign"
	mock "email-dispatch-gateway/internal/domain/campaign/mock"
	internalerrors "email-dispatch-gateway/internal/internal-errors"
	"errors"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
)

func Test_Service_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock.NewMockRepository(ctrl)
	service := campaign.NewService(repository)

	newCampaignDTO := contract.NewCampaignDTO{
		Name:    "Campaign Name",
		Content: "Campaign Content",
		Emails:  []string{"a@domain.com", "b@domain.com"},
	}

	t.Run("should create campaign", func(t *testing.T) {
		// ARRANGE
		repository.EXPECT().Save(gomock.Any()).Return(nil)

		// ACT
		id, err := service.Create(newCampaignDTO)

		// ASSERT
		require.NotEmpty(t, id)
		require.Empty(t, err)
	})

	t.Run("should return an error when a domain error occurs", func(t *testing.T) {
		// ARRANGE
		invalidDTO := newCampaignDTO
		invalidDTO.Name = ""

		// ACT
		id, err := service.Create(invalidDTO)

		// ASSERT
		require.Empty(t, id)
		require.Error(t, err)
	})

	t.Run("should call Repository.Save with correct arguments", func(t *testing.T) {
		// ARRANGE
		repository.EXPECT().Save(gomock.Cond(func(arguments any) bool {
			return arguments.(*campaign.Campaign).ID != "" &&
				arguments.(*campaign.Campaign).Name == newCampaignDTO.Name &&
				arguments.(*campaign.Campaign).Content == newCampaignDTO.Content &&
				len(arguments.(*campaign.Campaign).Contacts) == len(newCampaignDTO.Emails)
		})).Return(nil)

		// ACT
		id, err := service.Create(newCampaignDTO)

		// ASSERT
		require.NotEmpty(t, id)
		require.Empty(t, err)
	})

	t.Run("should return an internal error when a repository error occurs", func(t *testing.T) {
		// ARRANGE
		repository.EXPECT().Save(gomock.Any()).Return(errors.New("any repository error"))

		// ACT
		id, err := service.Create(newCampaignDTO)

		// ASSERT
		require.Empty(t, id)
		require.Equal(t, internalerrors.ErrInternalServerError, err)
	})
}
