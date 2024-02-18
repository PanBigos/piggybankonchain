package token

import (
	"errors"
	"time"

	"github.com/Exca-DK/pegism/core/clock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
)

// Different types of error returned by the VerifyToken function
var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// Payload contains the payload data of the token
type Payload struct {
	ID        uuid.UUID      `json:"id"`
	Address   common.Address `json:"username"`
	IssuedAt  int64          `json:"issued_at"`
	ExpiredAt int64          `json:"expired_at"`
	clock     *clock.Clock
}

// NewPayload creates a new token payload with a specific username and duration
func NewPayload(username string, duration time.Duration, clock *clock.Clock) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	now := clock.Now()
	payload := &Payload{
		ID:        tokenID,
		Address:   common.HexToAddress(username),
		IssuedAt:  now.Unix(),
		ExpiredAt: now.Add(duration).Unix(),
		clock:     clock,
	}
	return payload, nil
}

// Valid checks if the token payload is valid or not
func (payload *Payload) Valid() error {
	if payload.clock.Now().After(payload.ExpiredAtAsTime()) {
		return ErrExpiredToken
	}
	return nil
}

// Valid checks if the token payload is valid or not
func (payload *Payload) ExpiredAtAsTime() time.Time {
	return time.Unix(payload.ExpiredAt, 0)
}

// Valid checks if the token payload is valid or not
func (payload *Payload) IssuedAtAsTime() time.Time {
	return time.Unix(payload.IssuedAt, 0)
}
