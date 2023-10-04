package endpoints

import (
	"context"
	"errors"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Auth(t *testing.T) {
	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	t.Run("should return an error when authorization token is not provided", func(t *testing.T) {
		// ARRANGE
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()

		// ACT
		Auth(handlerFunc).ServeHTTP(w, r)

		// ASSERT
		require.Equal(t, w.Code, http.StatusUnauthorized)
		require.Equal(t, "{\"error\":\"request does not contain an authorization header\"}\n", w.Body.String())
	})

	t.Run("should return an error if oidc provider fail to generate a new provider", func(t *testing.T) {
		// ARRANGE
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/", nil)
		r.Header.Set("Authorization", "Bearer AnyToken")

		originalProvider := newOIDCProvider
		newOIDCProvider = func(ctx context.Context, realm string) (*oidc.Provider, error) {
			return nil, errors.New("any provider error")
		}
		defer func() { newOIDCProvider = originalProvider }()

		// ACT
		Auth(handlerFunc).ServeHTTP(w, r)

		// ASSERT
		require.Equal(t, w.Code, http.StatusInternalServerError)
		require.Equal(t, "{\"error\":\"internal server error\"}\n", w.Body.String())
	})
}
