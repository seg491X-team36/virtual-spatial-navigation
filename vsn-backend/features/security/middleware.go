package security

import (
	"context"
	"net/http"

	"github.com/seg491X-team36/vsn-backend/features/security/verification"
)

var tokenHeader = "token" // tokens should be stored in

type TokenVerifier interface {
	Verify(tokenString string) (verification.Token, error)
}

func AuthToken(ctx context.Context) (token verification.Token, ok bool) {
	token, ok = ctx.Value(tokenHeader).(verification.Token)
	return token, ok
}

func AuthMiddleware(verifier TokenVerifier) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// verify the token
			token, err := verifier.Verify(r.Header.Get(tokenHeader))

			// authentication required
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				_, _ = w.Write([]byte("authorization required"))
				return
			}

			// attach the token to the context
			ctx := context.WithValue(r.Context(), tokenHeader, token)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
