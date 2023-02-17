package experiment

import (
	"context"
	"net/http"

	"github.com/seg491X-team36/vsn-backend/domain/services"
)

type VerificationHandlers struct {
	VerificationService services.VerificationService
}

func (v *VerificationHandlers) EnterEmail() http.HandlerFunc {
	return postRequestWrapper(func(ctx context.Context, req verificationEmailRequest) verificationEmailResponse {
		// call the enter email method of the verification service
		v.VerificationService.EnterEmail(req.Email)
		return verificationEmailResponse{}
	})
}

func (v *VerificationHandlers) EnterVerificationCode() http.HandlerFunc {
	return postRequestWrapper(func(ctx context.Context, req verificationCodeRequest) verificationCodeResponse {
		// call the enter verification code method of the verification service
		token, err := v.VerificationService.EnterVerificationCode(req.Email, req.Code)

		return verificationCodeResponse{
			Token: token,
			Error: errWrapper(err),
		}
	})
}
