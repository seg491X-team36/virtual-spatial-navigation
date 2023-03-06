package verification

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type claimData[C any] struct {
	Claims C `json:"claims"`
	jwt.RegisteredClaims
}

type TokenManager[C any] struct {
	Secret       []byte
	ExpiresAfter time.Duration
}

// Verify a token
func (mgr *TokenManager[C]) Verify(token string) (claims C, err error) {
	data := &claimData[C]{}

	// parse the token into claims struct
	_, err = jwt.ParseWithClaims(token, data, func(raw *jwt.Token) (interface{}, error) {
		if _, ok := raw.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", raw.Header["alg"])
		}
		return mgr.Secret, nil
	})

	if err != nil {
		return claims, err
	}

	return data.Claims, nil
}

// Generate a token
func (mgr *TokenManager[C]) Generate(claims C) string {
	data := mgr.generate(claims)
	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, data).SignedString(mgr.Secret)
	return token
}

func (mgr *TokenManager[C]) generate(claims C) *claimData[C] {
	now := time.Now()

	return &claimData[C]{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(mgr.ExpiresAfter)),
		},
		Claims: claims,
	}
}
