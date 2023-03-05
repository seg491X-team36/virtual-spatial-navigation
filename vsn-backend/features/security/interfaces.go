package security

import (
	"context"

	"github.com/seg491X-team36/vsn-backend/domain/model"
)

const tokenHeader = "token" // tokens should be stored in

type tokenVerifier[C any] interface {
	Verify(token string) (C, error)
}

type UserVerifier tokenVerifier[model.UserClaims]

type AdminVerifier tokenVerifier[model.UserClaims]

func GetUserClaims(ctx context.Context) (claims model.UserClaims, ok bool) {
	claims, ok = ctx.Value(tokenHeader).(model.UserClaims)
	return claims, ok
}

func GetAdminClaims(ctx context.Context) (claims model.AdminClaims, ok bool) {
	claims, ok = ctx.Value(tokenHeader).(model.AdminClaims)
	return claims, ok
}
