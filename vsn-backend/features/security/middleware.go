package security

import (
	"context"
	"net/http"
)

func Middleware(verifier UserVerifier) func(http.Handler) http.Handler {
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
