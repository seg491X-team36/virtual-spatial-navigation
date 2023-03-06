package security

import (
	"errors"

	"github.com/seg491X-team36/vsn-backend/domain/model"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	Accounts       map[string][]byte
	TokenGenerator tokenGenerator[model.AdminClaims]
}

func (login *LoginService) Login(email, password string) (string, error) {
	hash, ok := login.Accounts[email]
	if !ok {
		// dummy operation to reduce difference in response time
		_ = bcrypt.CompareHashAndPassword(hash, []byte("dummy operation"))
		return "", errors.New("the email or password was incorrect")
	}

	// check the passwords match
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		return "", errors.New("the email or password was incorrect")
	}

	// generate the token
	token := login.TokenGenerator.Generate(model.AdminClaims{
		Email: email,
	})

	return token, nil
}
