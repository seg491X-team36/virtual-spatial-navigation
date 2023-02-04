package verification

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/seg491X-team36/vsn-backend/domain/model"
	"github.com/stretchr/testify/assert"
)

type verificationUserRepositoryStub struct {
	user model.User
	err  error
}

func (s *verificationUserRepositoryStub) GetByEmail(email string) (model.User, error) {
	return s.user, s.err
}

type verificationEmailServiceStub struct {
	sent map[string]string
}

func (s *verificationEmailServiceStub) Send(to, email string) {
	s.sent[to] = email
}

func TestService(t *testing.T) {
	email := &verificationEmailServiceStub{
		sent: map[string]string{},
	}

	service := Service{
		Users: &verificationUserRepositoryStub{
			user: model.User{Id: uuid.New()},
			err:  nil,
		},
		Email: email,
		Codes: &CodeGenerator{
			Characters: DefaultCharacters,
			Length:     6,
			Rand:       CryptoRandFunc,
		},
		CodesExpireAfter: time.Minute * 15,
		Tokens: &TokenManager{
			Secret: []byte("secret"),
		},
		Pending: map[string]verificationRequest{},
	}

	t.Run("happy-path", func(t *testing.T) {
		code, err := service.EnterEmail("test@gmail.com")
		assert.NoError(t, err)

		_, err = service.EnterVerificationCode("test@gmail.com", code)
		assert.NoError(t, err)

		// the verification request is deleted
		assert.Equal(t, 0, len(service.Pending))

		// the email has been sent
		_, ok := email.sent["test@gmail.com"]
		assert.True(t, ok)
	})

	t.Run("user-does-not-exist", func(t *testing.T) {
		service.Users = &verificationUserRepositoryStub{
			user: model.User{},
			err:  errors.New("user does not exist"),
		}

		_, err := service.EnterEmail("test@gmail.com")
		assert.Error(t, err)
	})

}

func TestVerify(t *testing.T) {
	service := Service{
		Tokens: &TokenManager{
			Secret: []byte("secret"),
		},
		Pending: map[string]verificationRequest{
			"test2@gmail.com": {
				code:     "111222",
				expireAt: time.Now().Add(-time.Minute), // expired
				userId:   uuid.UUID{},
			},
			"test3@gmail.com": {
				code:     "222333",
				expireAt: time.Now().Add(time.Minute), // not expired
				userId:   uuid.UUID{},
			},
		},
	}

	// does not match any email
	_, err := service.verify("test1@gmail.com", "")
	assert.Error(t, err)

	// does not match code
	_, err = service.verify("test2@gmail.com", "DOES NOT MATCH")
	assert.Error(t, err)

	// expired
	_, err = service.verify("test2@gmail.com", "111222")
	assert.Error(t, err)

	// successful verification
	_, err = service.verify("test3@gmail.com", "222333")
	assert.NoError(t, err)
}
