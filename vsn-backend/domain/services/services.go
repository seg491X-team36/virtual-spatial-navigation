package services

import (
	"context"
	"io"

	"github.com/seg491X-team36/vsn-backend/domain/model"
)

type FormWatcherService interface {
	Watch(ctx context.Context, onResponse func(response model.FormResponse))
}

type FormPollerService interface {
	Poll(ctx context.Context) []model.FormResponse
}

type EmailService interface {
	Send(to string, message string)
	SendHTML(to string, message io.Reader)
}

type VerificationService interface {
	EnterEmail(email string) error
	EnterVerificationCode(email, code string) (token string, err error)
}

type LoginService interface {
	Login(email, password string) (token string, err error)
}
