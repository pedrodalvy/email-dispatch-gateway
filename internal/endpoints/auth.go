package endpoints

import (
	internalErrors "email-dispatch-gateway/internal/internal-errors"
	"fmt"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/go-chi/render"
	"net/http"
	"strings"
)

const (
	aud   = "email-dispatch-gateway"
	realm = "email-dispatch-gateway"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{"error": "request does not contain an authorization header"})
			return
		}

		token = strings.Replace(token, "Bearer ", "", 1)

		provider, err := oidc.NewProvider(r.Context(), fmt.Sprintf("http://localhost:8080/realms/%s", realm))
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, map[string]string{"error": internalErrors.ErrInternalServerError.Error()})
			return
		}

		verifier := provider.Verifier(&oidc.Config{ClientID: aud})
		_, err = verifier.Verify(r.Context(), token)
		if err != nil {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{"error": "invalid token"})
			return
		}

		next.ServeHTTP(w, r)
	})
}
