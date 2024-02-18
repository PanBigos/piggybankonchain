// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package sqlc

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type Querier interface {
	AddMessage(ctx context.Context, arg AddMessageParams) (Message, error)
	CheckUserRegistration(ctx context.Context, address string) (bool, error)
	CreatePiggy(ctx context.Context, arg CreatePiggyParams) (Piggy, error)
	CreateProfile(ctx context.Context, address string) (Profile, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	GetMessageByHash(ctx context.Context, transactionHash string) (Message, error)
	GetMessagesByAddress(ctx context.Context, address string) ([]Message, error)
	GetPiggies(ctx context.Context, profileAddress string) ([]Piggy, error)
	GetPiggy(ctx context.Context, address string) (Piggy, error)
	GetPiggyFromName(ctx context.Context, name sql.NullString) (Piggy, error)
	GetPiggyFromProfileAddress(ctx context.Context, profileAddress string) (Piggy, error)
	GetProfile(ctx context.Context, address string) (Profile, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	UpdatePiggy(ctx context.Context, arg UpdatePiggyParams) (Piggy, error)
}

var _ Querier = (*Queries)(nil)
