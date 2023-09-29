package endpoints

import (
	internalErrors "email-dispatch-gateway/internal/internal-errors"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_HandlerError(t *testing.T) {
	t.Run("should return an internal error", func(t *testing.T) {
		// ARRANGE
		endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
			return nil, http.StatusOK, internalErrors.ErrInternalServerError
		}

		handlerFunc := HandlerError(endpoint)
		req, _ := http.NewRequest("GET", "/", nil)
		res := httptest.NewRecorder()

		// ACT
		handlerFunc.ServeHTTP(res, req)

		// ASSERT
		require.Equal(t, http.StatusInternalServerError, res.Code)
		require.Contains(t, res.Body.String(), internalErrors.ErrInternalServerError.Error())
	})

	t.Run("should return a domain error", func(t *testing.T) {
		// ARRANGE
		endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
			return nil, http.StatusOK, errors.New("any domain error")
		}

		handlerFunc := HandlerError(endpoint)
		req, _ := http.NewRequest("GET", "/", nil)
		res := httptest.NewRecorder()

		// ACT
		handlerFunc.ServeHTTP(res, req)

		// ASSERT
		require.Equal(t, http.StatusBadRequest, res.Code)
		require.Contains(t, res.Body.String(), "any domain error")
	})

	t.Run("should return a success response", func(t *testing.T) {
		// ARRANGE
		type bodyForTest struct{ ID int }
		expectedResponseData := bodyForTest{ID: 1}

		endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
			return expectedResponseData, http.StatusCreated, nil
		}

		handlerFunc := HandlerError(endpoint)
		req, _ := http.NewRequest("POST", "/", nil)
		res := httptest.NewRecorder()

		// ACT
		handlerFunc.ServeHTTP(res, req)

		// ASSERT
		var returnedData bodyForTest
		json.Unmarshal(res.Body.Bytes(), &returnedData)

		require.Equal(t, expectedResponseData, returnedData)
		require.Equal(t, http.StatusCreated, res.Code)
	})

}
