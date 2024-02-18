package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Exca-DK/pegism/core/clock"
	"github.com/Exca-DK/pegism/core/database"
	"github.com/Exca-DK/pegism/core/token"
	c_types "github.com/Exca-DK/pegism/core/types"
	"github.com/Exca-DK/pegism/service/types"
	"github.com/google/uuid"

	"github.com/Exca-DK/pegism/core/log"
	"github.com/Exca-DK/pegism/service/db/sqlc"

	"github.com/ethereum/go-ethereum/common"
	_ "github.com/lib/pq"
)

type SessionDatabase interface {
	CreateSession(
		ctx context.Context,
		signedToken token.Token,
		token *token.Payload,
	) error
	GetSession(
		ctx context.Context,
		id uuid.UUID,
	) (c_types.Session, error)
}

type ProfileDatabase interface {
	CreateProfile(
		context.Context,
		common.Address,
	) (types.Profile, error)
	GetProfile(context.Context, common.Address) (types.Profile, error)
	Registered(context.Context, common.Address) (bool, error)
}

type PiggyDatabase interface {
	CreatePiggy(
		context.Context,
		types.Piggy,
	) (types.Piggy, error)
	GetPiggy(context.Context, common.Address) (types.Piggy, error)
	GetPiggyFromName(context.Context, string) (types.Piggy, error)
	GetPiggyFromProfile(context.Context, common.Address) (types.Piggy, error)
	UpdatePiggyName(context.Context, common.Address, string) (types.Piggy, error)
	AddMessage(context.Context, types.Message) (types.Message, error)
}

type Store interface {
	ProfileDatabase
	SessionDatabase
	PiggyDatabase
}

var _ Store = (*Database)(nil)

type Database struct {
	sqlc        *sqlc.Queries
	db          *sql.DB
	inner       database.PostgresStore
	interceptor database.StoreErrorInterceptor
	log         log.Logger

	clock *clock.Clock
}

type Config struct {
	Username string
	Password string
	Db       string
	Endpoint string
	Clock    *clock.Clock
}

func (cfg Config) validate() error {
	if cfg.Clock == nil {
		return errors.New("clock not provided for db config")
	}
	return nil
}

func NewStore(cfg Config) (*Database, error) {
	if err := cfg.validate(); err != nil {
		return nil, err
	}
	conn, err := database.NewPostgresUnsecureConn(database.Config{
		Username: cfg.Username,
		Password: cfg.Password,
		Db:       cfg.Db,
		Endpoint: cfg.Endpoint,
	})
	if err != nil {
		return nil, fmt.Errorf("NewStoreFailure: %w", err)
	}
	store := &Database{
		sqlc:  sqlc.New(conn),
		db:    conn,
		inner: database.PostgresStore{Db: conn},
		log:   log.Root().With("db", "pegism"),
		clock: cfg.Clock,
	}
	return store, nil
}

func (store *Database) CreateProfile(
	ctx context.Context,
	address common.Address,
) (types.Profile, error) {

	profile, err := store.sqlc.CreateProfile(ctx, address.Hex())
	if err != nil {
		return types.Profile{}, store.interceptor.InterceptError(err)
	}
	adapter := profileAdapter{q: store.sqlc}
	return adapter.FromModel(profile)
}

func (store *Database) GetProfile(
	ctx context.Context,
	streamer common.Address,
) (types.Profile, error) {
	model, err := store.sqlc.GetProfile(ctx, streamer.Hex())
	if err != nil {
		return types.Profile{}, store.interceptor.InterceptError(err)
	}
	adapter := profileAdapter{q: store.sqlc}
	return adapter.FromModel(model)
}

func (store *Database) Registered(
	ctx context.Context,
	streamer common.Address,
) (bool, error) {
	return store.sqlc.CheckUserRegistration(ctx, streamer.Hex())
}

func (store *Database) CreateSession(
	ctx context.Context,
	signedToken token.Token,
	token *token.Payload,
) error {
	_, err := store.sqlc.CreateSession(ctx, sqlc.CreateSessionParams{
		ID:           token.ID,
		Address:      token.Address.Hex(),
		RefreshToken: string(signedToken),
		UserAgent:    "",
		ClientIp:     "",
		ExpiresAt:    token.ExpiredAtAsTime(),
	})
	return store.interceptor.InterceptError(err)
}

func (store *Database) GetSession(
	ctx context.Context,
	id uuid.UUID,
) (c_types.Session, error) {
	obj, err := store.sqlc.GetSession(ctx, id)
	if err != nil {
		return c_types.Session{}, store.interceptor.InterceptError(err)
	}

	return c_types.Session{
		Address:      common.HexToAddress(obj.Address),
		Id:           id,
		RefreshToken: obj.RefreshToken,
		IsBlocked:    obj.IsBlocked,
		CreatedAt:    obj.CreatedAt,
		ExpiresAt:    obj.ExpiresAt,
	}, nil
}

func (store *Database) CreatePiggy(
	ctx context.Context,
	args types.Piggy,
) (types.Piggy, error) {
	adapter := piggyAdapter{}
	model := adapter.ToModel(args)
	created, err := store.sqlc.CreatePiggy(ctx, sqlc.CreatePiggyParams{
		Address:        model.Address,
		FromAddress:    model.FromAddress,
		ProfileAddress: model.ProfileAddress,
		CreatedAt:      model.CreatedAt,
		UnlocksAt:      model.UnlocksAt,
		Name:           model.Name,
	})
	if err != nil {
		return types.Piggy{}, store.interceptor.InterceptError(err)
	}
	return adapter.FromModel(created, nil), nil
}

func (store *Database) GetPiggy(ctx context.Context, address common.Address) (types.Piggy, error) {
	piggyModel, err := store.sqlc.GetPiggy(ctx, address.Hex())
	if err != nil {
		return types.Piggy{}, store.interceptor.InterceptError(err)
	}
	piggyMessages, err := store.sqlc.GetMessagesByAddress(ctx, piggyModel.Address)
	if err != nil {
		return types.Piggy{}, store.interceptor.InterceptError(err)
	}
	p_adapter := piggyAdapter{}
	adaptedPiggy := p_adapter.FromModel(piggyModel, piggyMessages)
	return adaptedPiggy, nil
}

func (store *Database) GetPiggyFromProfile(
	ctx context.Context,
	address common.Address,
) (types.Piggy, error) {
	piggyModel, err := store.sqlc.GetPiggyFromProfileAddress(ctx, address.Hex())
	if err != nil {
		return types.Piggy{}, store.interceptor.InterceptError(err)
	}
	piggyMessages, err := store.sqlc.GetMessagesByAddress(ctx, piggyModel.Address)
	if err != nil {
		return types.Piggy{}, store.interceptor.InterceptError(err)
	}
	p_adapter := piggyAdapter{}
	adaptedPiggy := p_adapter.FromModel(piggyModel, piggyMessages)
	return adaptedPiggy, nil
}

func (store *Database) GetPiggyFromName(ctx context.Context, name string) (types.Piggy, error) {
	query := sql.NullString{String: name, Valid: true}
	piggyModel, err := store.sqlc.GetPiggyFromName(ctx, query)
	if err != nil {
		return types.Piggy{}, store.interceptor.InterceptError(err)
	}
	piggyMessages, err := store.sqlc.GetMessagesByAddress(ctx, piggyModel.Address)
	if err != nil {
		return types.Piggy{}, store.interceptor.InterceptError(err)
	}
	p_adapter := piggyAdapter{}
	adaptedPiggy := p_adapter.FromModel(piggyModel, piggyMessages)
	return adaptedPiggy, nil
}

func (store *Database) UpdatePiggyName(
	ctx context.Context,
	address common.Address,
	name string,
) (types.Piggy, error) {
	piggyModel, err := store.sqlc.UpdatePiggy(ctx, sqlc.UpdatePiggyParams{
		Name:    sql.NullString{String: name, Valid: true},
		Address: address.Hex(),
	})
	if err != nil {
		return types.Piggy{}, store.interceptor.InterceptError(err)
	}
	piggyMessages, err := store.sqlc.GetMessagesByAddress(ctx, address.Hex())
	if err != nil {
		return types.Piggy{}, store.interceptor.InterceptError(err)
	}
	p_adapter := piggyAdapter{}
	adaptedPiggy := p_adapter.FromModel(piggyModel, piggyMessages)
	return adaptedPiggy, nil
}

func (store *Database) AddMessage(ctx context.Context, msg types.Message) (types.Message, error) {
	adapter := messageAdapter{}
	adapted := adapter.ToModel(msg)
	piggyModel, err := store.sqlc.AddMessage(ctx, sqlc.AddMessageParams{
		TransactionHash: adapted.TransactionHash,
		Address:         adapted.Address,
		Token:           adapted.Token,
		Amount:          adapted.Amount,
		Fee:             adapted.Fee,
		Content:         adapted.Content,
		Nick:            adapted.Nick,
	})
	if err != nil {
		return types.Message{}, store.interceptor.InterceptError(err)
	}
	return adapter.FromModel(piggyModel), nil
}
