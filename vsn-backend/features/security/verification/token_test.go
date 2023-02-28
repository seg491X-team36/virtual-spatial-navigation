package verification

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTokenManager(t *testing.T) {
	mgr := &TokenManager{
		Secret: []byte("secret-1"),
	}

	userId := uuid.New()

	// generating a token
	tokenString := mgr.Generate(Token{UserId: userId})

	// verifying a token
	token, err := mgr.Verify(tokenString)
	assert.NoError(t, err)
	assert.Equal(t, userId, token.UserId) // user IDs must match

	// change secret and verification fails
	mgr.Secret = []byte("secret-2")
	_, err = mgr.Verify(tokenString)
	assert.Error(t, err)
}
