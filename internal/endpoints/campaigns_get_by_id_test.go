package endpoints

import (
	"context"
	"email-dispatch-gateway/internal/contract"
	"email-dispatch-gateway/internal/domain/campaign"
	mock "email-dispatch-gateway/internal/domain/campaign/mock"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Handler_CampaignsGetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mock.NewMockServiceInterface(ctrl)

	t.Run("should return a campaign", func(t *testing.T) {
		// ARRANGE
		campaignResponse := contract.CampaignResponse{
			ID:      "any id",
			Name:    "any name",
			Content: "any content",
			Status:  campaign.Pending,
		}

		service.EXPECT().GetByID(gomock.Eq(campaignResponse.ID)).Return(campaignResponse, nil)
		handler := NewHandler(service)

		res := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/", nil)

		chiContext := chi.NewRouteContext()
		chiContext.URLParams.Add("id", campaignResponse.ID)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiContext))

		// ACT
		responseData, status, err := handler.CampaignsGetByID(res, req)

		// ASSERT
		require.Equal(t, campaignResponse, responseData)
		require.Equal(t, http.StatusOK, status)
		require.Nil(t, err)
	})

	t.Run("should return an error when a domain error occurs", func(t *testing.T) {
		// ARRANGE
		service.EXPECT().GetByID(gomock.Any()).Return(contract.CampaignResponse{}, errors.New("any domain error"))
		handler := NewHandler(service)

		res := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/", nil)

		chiContext := chi.NewRouteContext()
		chiContext.URLParams.Add("id", "campaignID")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiContext))

		// ACT
		responseData, status, err := handler.CampaignsGetByID(res, req)

		// ASSERT
		require.Empty(t, responseData)
		require.Equal(t, http.StatusOK, status)
		require.Equal(t, "any domain error", err.Error())
	})
}
