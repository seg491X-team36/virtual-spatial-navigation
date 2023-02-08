package verification

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type claims struct {
	Token `json:"claims"`
	jwt.RegisteredClaims
}

type Token struct {
	UserId uuid.UUID `json:"userId"`
}

type TokenManager struct {
	Secret []byte
}

func (mgr *TokenManager) Verify(tokenString string) (Token, error) {
	c := &claims{}

	// parse the token into claims struct
	_, err := jwt.ParseWithClaims(tokenString, c, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return mgr.Secret, nil
	})

	if err != nil {
		return Token{}, err
	}

	return c.Token, nil
}

func (mgr *TokenManager) Generate(token Token) string {
	now := time.Now()

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   token.UserId.String(),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * 24)),
		},
		Token: token,
	}).SignedString(mgr.Secret)

	if err != nil {
		panic(err)
	}

	return tokenString
}
