package endpoints

import (
	"context"
	"net/http"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/golang-jwt/jwt/v5"
)

func CheckAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		token := strings.Split(authorization, " ")

		if len(token) < 2 || token[1] == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		provider, err := oidc.NewProvider(r.Context(), "http://localhost:8080/realms/provider")

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		verifier := provider.Verifier(&oidc.Config{ClientID: "emailn"})

		_, err = verifier.Verify(r.Context(), token[1])

		if err != nil {
			http.Error(w, "Invalid Token", http.StatusInternalServerError)
			return
		}

		parsedToken, _ := jwt.Parse(token[1], nil)
		claims := parsedToken.Claims.(jwt.MapClaims)
		email := claims["email"]

		ctx := context.WithValue(r.Context(), "email", email)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
