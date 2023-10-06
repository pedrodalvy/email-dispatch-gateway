package endpoints

import (
	"context"
	internalErrors "email-dispatch-gateway/internal/internal-errors"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/go-chi/render"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strings"
)

var newOIDCProvider = oidc.NewProvider

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		if tokenString == "" {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{"error": "request does not contain an authorization header"})
			return
		}

		provider, err := newOIDCProvider(r.Context(), os.Getenv("KEYCLOAK_ISSUER"))
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, map[string]string{"error": internalErrors.ErrInternalServerError.Error()})
			return
		}

		verifier := provider.Verifier(&oidc.Config{
			ClientID:        os.Getenv("KEYCLOAK_AUD"),
			SkipIssuerCheck: os.Getenv("KEYCLOAK_SKIP_ISSUER_CHECK") == "true",
		})
		_, err = verifier.Verify(r.Context(), tokenString)
		if err != nil {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{"error": "invalid token"})
			return
		}

		email := extractUserEmail(tokenString)
		ctx := context.WithValue(r.Context(), "email", email)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func extractUserEmail(tokenString string) string {
	parsedToken, _ := jwt.Parse(tokenString, nil)
	claims := parsedToken.Claims.(jwt.MapClaims)

	return claims["email"].(string)
}
