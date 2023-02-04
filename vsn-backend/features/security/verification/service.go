package verification

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/seg491X-team36/vsn-backend/domain/model"
)

type verificationUserRepository interface {
	GetByEmail(email string) (model.User, error)
}

type verificationEmailService interface {
	Send(to, message string)
}

type verificationRequest struct {
	code     string
	expireAt time.Time
	userId   uuid.UUID // saved during EnterEmail, used in EnterVerificationCode
}

type Service struct {
	Users            verificationUserRepository
	Email            verificationEmailService
	Tokens           *TokenManager
	Codes            *CodeGenerator
	CodesExpireAfter time.Duration
	Pending          map[string]verificationRequest
}

func (s *Service) EnterEmail(email string) (string, error) {
	// user must exist
	user, err := s.Users.GetByEmail(email)
	if err != nil {
		return "", errors.New("user does not exist")
	}

	code := s.Codes.Generate()
	message := fmt.Sprintf("verification code: %s", code)
	s.Email.Send(email, message)

	// save verification request
	s.Pending[email] = verificationRequest{
		code:     code,
		expireAt: time.Now().Add(s.CodesExpireAfter),
		userId:   user.Id,
	}

	return code, nil
}

func (s *Service) EnterVerificationCode(email, code string) (string, error) {
	// verify the user
	request, err := s.verify(email, code)
	if err != nil {
		return "", nil
	}
	delete(s.Pending, email)

	// create the token
	token := s.Tokens.Generate(Token{UserId: request.userId})
	return token, nil
}

func (s *Service) verify(email, code string) (verificationRequest, error) {
	request, ok := s.Pending[email]

	// verification code has not been requested
	if !ok {
		return verificationRequest{}, errors.New("verification failed")
	}

	// verification codes do not match
	if code != request.code {
		return verificationRequest{}, errors.New("verification failed")
	}

	// verification code has expired
	if time.Now().After(request.expireAt) {
		return verificationRequest{}, errors.New("verification failed")
	}

	return request, nil
}
