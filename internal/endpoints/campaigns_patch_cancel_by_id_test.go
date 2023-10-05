package endpoints

import (
	"context"
	mock "email-dispatch-gateway/internal/domain/campaign/mock"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Handler_CampaignsPatchCancelByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mock.NewMockServiceInterface(ctrl)

	t.Run("should cancel a campaign", func(t *testing.T) {
		// ARRANGE
		campaignID := "any"

		service.EXPECT().CancelByID(gomock.Eq(campaignID)).Return(nil)
		handler := NewHandler(service)

		res := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/", nil)

		chiContext := chi.NewRouteContext()
		chiContext.URLParams.Add("id", campaignID)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiContext))

		// ACT
		responseData, status, err := handler.CampaignsPatchCancelByID(res, req)

		// ASSERT
		require.Nil(t, responseData)
		require.Equal(t, http.StatusNoContent, status)
		require.Nil(t, err)
	})

	t.Run("should return an error when the service returns an error", func(t *testing.T) {
		// ARRANGE
		service.EXPECT().CancelByID(gomock.Any()).Return(errors.New("any service error"))
		handler := NewHandler(service)

		res := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/", nil)

		chiContext := chi.NewRouteContext()
		chiContext.URLParams.Add("id", "campaignID")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiContext))

		// ACT
		responseData, status, err := handler.CampaignsPatchCancelByID(res, req)

		// ASSERT
		require.Nil(t, responseData)
		require.Equal(t, http.StatusNoContent, status)
		require.Equal(t, "any service error", err.Error())
	})
}
