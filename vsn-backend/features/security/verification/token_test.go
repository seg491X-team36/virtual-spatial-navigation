package verification

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/seg491X-team36/vsn-backend/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestTokenManager(t *testing.T) {
	mgr := &TokenManager[model.UserClaims]{
		Secret:       []byte("secret-1"),
		ExpiresAfter: time.Minute * 1,
	}

	userId := uuid.New()

	// generating a token
	token := mgr.Generate(model.UserClaims{UserId: userId})

	// verifying a token
	claims, err := mgr.Verify(token)
	assert.NoError(t, err)
	assert.Equal(t, userId, claims.UserId) // user IDs must match

	// change secret and verification fails
	mgr.Secret = []byte("secret-2")
	_, err = mgr.Verify(token)
	assert.Error(t, err)
}
