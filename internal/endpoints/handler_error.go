package endpoints

import (
	internalErrors "email-dispatch-gateway/internal/internal-errors"
	"errors"
	"github.com/go-chi/render"
	"net/http"
)

type EndpointFunc func(w http.ResponseWriter, r *http.Request) (responseData interface{}, status int, err error)

func HandlerError(ef EndpointFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responseData, status, err := ef(w, r)

		if err != nil {
			status = getErrorCode(err)

			render.Status(r, status)
			render.JSON(w, r, map[string]string{"error": err.Error()})

			return
		}

		render.Status(r, status)
		render.JSON(w, r, responseData)
	}
}

func getErrorCode(err error) int {
	if errors.Is(err, internalErrors.ErrInternalServerError) {
		return http.StatusInternalServerError
	}

	if errors.Is(err, internalErrors.ErrResourceNotFound) {
		return http.StatusNotFound
	}

	return http.StatusBadRequest
}
