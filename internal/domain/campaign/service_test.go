package campaign_test

import (
	"email-dispatch-gateway/internal/contract"
	"email-dispatch-gateway/internal/domain/campaign"
	mock "email-dispatch-gateway/internal/domain/campaign/mock"
	internalErrors "email-dispatch-gateway/internal/internal-errors"
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
		Name:      "Campaign Name",
		Content:   "Campaign Content",
		Emails:    []string{"a@domain.com", "b@domain.com"},
		CreatedBy: "test@email.com",
	}

	t.Run("should create campaign", func(t *testing.T) {
		// ARRANGE
		repository.EXPECT().Create(gomock.Any()).Return(nil)

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
		require.NotEqual(t, internalErrors.ErrInternalServerError, err)
	})

	t.Run("should call Repository.Create with correct arguments", func(t *testing.T) {
		// ARRANGE
		repository.EXPECT().Create(gomock.Cond(func(arguments any) bool {
			return arguments.(*campaign.Campaign).ID != "" &&
				arguments.(*campaign.Campaign).Name == newCampaignDTO.Name &&
				arguments.(*campaign.Campaign).Content == newCampaignDTO.Content &&
				arguments.(*campaign.Campaign).CreatedBy == newCampaignDTO.CreatedBy &&
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
		repository.EXPECT().Create(gomock.Any()).Return(errors.New("any repository error"))

		// ACT
		id, err := service.Create(newCampaignDTO)

		// ASSERT
		require.Empty(t, id)
		require.Equal(t, internalErrors.ErrInternalServerError, err)
	})
}

func Test_Service_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock.NewMockRepository(ctrl)
	service := campaign.NewService(repository)

	t.Run("should return a campaign", func(t *testing.T) {
		// ARRANGE
		c, _ := campaign.NewCampaign("Campaign Name", "Campaign Content", []string{"a@domain.com"}, "test@email.com")
		expectedCampaign := contract.CampaignResponse{
			ID:                   c.ID,
			Name:                 c.Name,
			Content:              c.Content,
			Status:               c.Status,
			AmountOfEmailsToSend: len(c.Contacts),
		}

		repository.EXPECT().GetByID(gomock.Eq(c.ID)).Return(c, nil)

		// ACT
		receivedCampaign, err := service.GetByID(c.ID)

		// ASSERT
		require.Equal(t, expectedCampaign, receivedCampaign)
		require.Nil(t, err)
	})

	t.Run("should return an internal server error if repository returns an error", func(t *testing.T) {
		// ARRANGE
		repository.EXPECT().GetByID(gomock.Any()).Return(nil, errors.New("any repository error"))

		// ACT
		receivedCampaign, err := service.GetByID("any")

		// ASSERT
		require.Empty(t, receivedCampaign)
		require.Equal(t, internalErrors.ErrInternalServerError, err)
	})

	t.Run("should return a resource not found error if campaign does not exist", func(t *testing.T) {
		// ARRANGE
		repository.EXPECT().GetByID(gomock.Any()).Return(nil, nil)

		// ACT
		receivedCampaign, err := service.GetByID("any")

		// ASSERT
		require.Empty(t, receivedCampaign)
		require.Equal(t, internalErrors.ErrResourceNotFound, err)
	})
}

func Test_Service_CancelByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock.NewMockRepository(ctrl)
	service := campaign.NewService(repository)

	t.Run("should cancel a campaign", func(t *testing.T) {
		// ARRANGE
		c, _ := campaign.NewCampaign("Campaign Name", "Campaign Content", []string{"a@domain.com"}, "test@email.com")
		repository.EXPECT().GetByID(gomock.Eq(c.ID)).Return(c, nil)
		repository.EXPECT().Update(gomock.Eq(c)).Return(nil)

		// ACT
		err := service.CancelByID(c.ID)

		// ASSERT
		require.Nil(t, err)
	})

	t.Run("should return an internal server error if repository.GetByID returns an error", func(t *testing.T) {
		// ARRANGE
		campaignID := "any"
		repository.EXPECT().GetByID(gomock.Eq(campaignID)).Return(&campaign.Campaign{}, errors.New("any repository error"))

		// ACT
		err := service.CancelByID(campaignID)

		// ASSERT
		require.Equal(t, internalErrors.ErrInternalServerError, err)
	})

	t.Run("should return a domain error", func(t *testing.T) {
		// ARRANGE
		c, _ := campaign.NewCampaign("Campaign Name", "Campaign Content", []string{"a@domain.com"}, "test@email.com")
		c.Cancel()
		repository.EXPECT().GetByID(gomock.Eq(c.ID)).Return(c, nil)

		// ACT
		err := service.CancelByID(c.ID)

		// ASSERT
		require.Error(t, err)
		require.NotEqual(t, internalErrors.ErrInternalServerError, err)
	})

	t.Run("should return an internal server error if repository.Update returns an error", func(t *testing.T) {
		// ARRANGE
		c, _ := campaign.NewCampaign("Campaign Name", "Campaign Content", []string{"a@domain.com"}, "test@email.com")
		repository.EXPECT().GetByID(gomock.Eq(c.ID)).Return(c, nil)
		repository.EXPECT().Update(gomock.Eq(c)).Return(errors.New("any repository error"))

		// ACT
		err := service.CancelByID(c.ID)

		// ASSERT
		require.Equal(t, internalErrors.ErrInternalServerError, err)
	})

	t.Run("should return a resource not found error if campaign does not exist", func(t *testing.T) {
		// ARRANGE
		repository.EXPECT().GetByID(gomock.Any()).Return(nil, nil)

		// ACT
		err := service.CancelByID("any")

		// ASSERT
		require.Equal(t, internalErrors.ErrResourceNotFound, err)
	})
}

func Test_Service_DeleteByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock.NewMockRepository(ctrl)
	service := campaign.NewService(repository)

	t.Run("should delete a campaign", func(t *testing.T) {
		// ARRANGE
		c, _ := campaign.NewCampaign("Campaign Name", "Campaign Content", []string{"a@domain.com"}, "test@email.com")
		repository.EXPECT().GetByID(gomock.Eq(c.ID)).Return(c, nil)
		repository.EXPECT().Delete(gomock.Eq(c)).Return(nil)

		// ACT
		err := service.DeleteByID(c.ID)

		// ASSERT
		require.Nil(t, err)
	})

	t.Run("should return an internal server error if repository.GetByID returns an error", func(t *testing.T) {
		// ARRANGE
		campaignID := "any"
		repository.EXPECT().GetByID(gomock.Eq(campaignID)).Return(&campaign.Campaign{}, errors.New("any repository error"))

		// ACT
		err := service.DeleteByID(campaignID)

		// ASSERT
		require.Equal(t, internalErrors.ErrInternalServerError, err)
	})

	t.Run("should return a domain error", func(t *testing.T) {
		// ARRANGE
		c, _ := campaign.NewCampaign("Campaign Name", "Campaign Content", []string{"a@domain.com"}, "test@email.com")
		c.Status = "another"
		repository.EXPECT().GetByID(gomock.Eq(c.ID)).Return(c, nil)

		// ACT
		err := service.DeleteByID(c.ID)

		// ASSERT
		require.Error(t, err)
		require.NotEqual(t, internalErrors.ErrInternalServerError, err)
	})

	t.Run("should return an internal server error if repository.Delete returns an error", func(t *testing.T) {
		// ARRANGE
		c, _ := campaign.NewCampaign("Campaign Name", "Campaign Content", []string{"a@domain.com"}, "test@email.com")
		repository.EXPECT().GetByID(gomock.Eq(c.ID)).Return(c, nil)
		repository.EXPECT().Delete(gomock.Eq(c)).Return(errors.New("any repository error"))

		// ACT
		err := service.DeleteByID(c.ID)

		// ASSERT
		require.Equal(t, internalErrors.ErrInternalServerError, err)
	})

	t.Run("should return a resource not found error if campaign does not exist", func(t *testing.T) {
		// ARRANGE
		repository.EXPECT().GetByID(gomock.Any()).Return(nil, nil)

		// ACT
		err := service.DeleteByID("any")

		// ASSERT
		require.Equal(t, internalErrors.ErrResourceNotFound, err)
	})
}
