package security

import (
	"testing"
	"time"

	"github.com/seg491X-team36/vsn-backend/domain/model"
	"github.com/seg491X-team36/vsn-backend/features/security/verification"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestLoginService(t *testing.T) {
	hash, _ := bcrypt.GenerateFromPassword([]byte("password"), 5)

	service := &LoginService{
		Accounts: map[string][]byte{
			"test@test.com": hash,
		},
		TokenGenerator: &verification.TokenManager[model.AdminClaims]{
			Secret:       []byte("secret"),
			ExpiresAfter: 5 * time.Minute,
		},
	}

	// wrong/non-existent account
	_, err := service.Login("test2@test.com", "password")
	assert.Error(t, err)

	// wrong password
	_, err = service.Login("test@test.com", "1234")
	assert.Error(t, err)

	// correct
	_, err = service.Login("test@test.com", "password")
	assert.NoError(t, err)

}
