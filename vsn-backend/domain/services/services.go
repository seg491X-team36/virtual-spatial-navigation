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
	EnterEmail(email string) (nonce int)
	EnterVerificationCode(email, code string, nonce int) (token string, err error)
}
