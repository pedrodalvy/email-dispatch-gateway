package endpoints

import (
	"bytes"
	"email-dispatch-gateway/internal/contract"
	mock "email-dispatch-gateway/internal/domain/campaign/mock"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_TestHandler_CampaignsPost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	campaignDTO := contract.NewCampaignDTO{
		Name:    "Test Name",
		Content: "Test Content",
		Emails:  []string{"test@domain.com"},
	}
	var dtoBuffer bytes.Buffer
	json.NewEncoder(&dtoBuffer).Encode(campaignDTO)

	service := mock.NewMockServiceInterface(ctrl)

	t.Run("should create a new campaign", func(t *testing.T) {
		// ARRANGE
		campaignID := "123"

		service.EXPECT().Create(gomock.Eq(campaignDTO)).Return(campaignID, nil)
		handler := NewHandler(service)

		req, _ := http.NewRequest("POST", "/", &dtoBuffer)
		res := httptest.NewRecorder()

		// ACT
		responseData, status, err := handler.CampaignsPost(res, req)

		// ASSERT
		require.Equal(t, map[string]string{"id": campaignID}, responseData)
		require.Equal(t, http.StatusCreated, status)
		require.Nil(t, err)
	})

	t.Run("should return an error when a domain error occurs", func(t *testing.T) {
		// ARRANGE
		service.EXPECT().Create(gomock.Any()).Return("", errors.New("any domain error"))
		handler := NewHandler(service)

		req, _ := http.NewRequest("POST", "/", &dtoBuffer)
		res := httptest.NewRecorder()

		// ACT
		responseData, status, err := handler.CampaignsPost(res, req)

		// ASSERT
		require.Empty(t, responseData)
		require.Equal(t, http.StatusCreated, status)
		require.Equal(t, "any domain error", err.Error())
	})
}
