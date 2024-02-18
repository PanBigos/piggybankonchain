package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/Exca-DK/pegism/core/clock"
	"github.com/dgrijalva/jwt-go"
)

type TokenFactory interface {
	CreateToken(username string, duration time.Duration) (Token, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}

type Token string

const minSecretKeySize = 32

type jwtFactory struct {
	secretKey string
	clock     *clock.Clock
}

func NewJwtFactory(secretKey string, clock *clock.Clock) (*jwtFactory, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf(
			"invalid token secret size: must be at least %d characters",
			minSecretKeySize,
		)
	}
	return &jwtFactory{secretKey: secretKey, clock: clock}, nil
}

// CreateToken creates a new token for a specific username and duration
func (maker *jwtFactory) CreateToken(
	username string,
	duration time.Duration,
) (Token, *Payload, error) {
	payload, err := NewPayload(username, duration, maker.clock)
	if err != nil {
		return "", payload, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(maker.secretKey))
	return Token(token), payload, err
}

// VerifyToken checks if the token is valid or not
func (maker *jwtFactory) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}
	payload := &Payload{clock: maker.clock}
	jwtToken, err := jwt.ParseWithClaims(token, payload, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}
	return payload, nil
}
