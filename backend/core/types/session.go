package types

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
)

type Session struct {
	Address      common.Address
	Id           uuid.UUID
	RefreshToken string
	IsBlocked    bool
	CreatedAt    time.Time
	ExpiresAt    time.Time
}
