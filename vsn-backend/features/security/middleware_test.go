package security

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/seg491X-team36/vsn-backend/features/security/verification"
	"github.com/stretchr/testify/assert"
)

type tokenVerifierStub struct {
	Expected string
	Token    verification.Token
}

func (t *tokenVerifierStub) Verify(token string) (verification.Token, error) {
	if token != t.Expected {
		return verification.Token{}, errors.New("token stub error")
	}
	return t.Token, nil
}

func TestMiddleware(t *testing.T) {
	id := uuid.New()

	verifier := &tokenVerifierStub{
		Expected: "EXPECTED",
		Token:    verification.Token{UserId: id},
	}

	middleware := AuthMiddleware(verifier)

	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get the token
		token, ok := AuthToken(r.Context())
		assert.Equal(t, id, token.UserId)
		assert.True(t, ok)
	}))

	// token is valid
	t.Run("happy-path", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
		r.Header.Add(tokenHeader, "EXPECTED")
		w := httptest.NewRecorder()

		// handle http request
		handler.ServeHTTP(w, r)
		res := w.Result()

		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

	// token is invalid
	t.Run("invalid-token", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
		r.Header.Add(tokenHeader, "NOT-EXPECTED")
		w := httptest.NewRecorder()

		// handle http request
		handler.ServeHTTP(w, r)
		res := w.Result()

		assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
	})

	// token is missing
	t.Run("missing-token", func(t *testing.T) {
		// do not attach a token to the request
		r := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
		w := httptest.NewRecorder()

		// handle http request
		handler.ServeHTTP(w, r)
		res := w.Result()

		assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
	})
}
