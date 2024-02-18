package backend

import (
	"errors"

	"github.com/google/uuid"

	"github.com/Exca-DK/pegism/core/token"
	c_types "github.com/Exca-DK/pegism/core/types"

	"github.com/Exca-DK/pegism/service/types"
	"github.com/ethereum/go-ethereum/common"
)

type SessionApi interface {
	CreateToken(address common.Address) (struct {
		AccessToken    token.Token
		AccessPayload  *token.Payload
		RefreshToken   token.Token
		RefreshPayload *token.Payload
	}, error)
	Renew(RefreshToken token.Token) (token.Token, *token.Payload, error)
	VerifyToken(string) (*token.Payload, error)
	GetSession(id uuid.UUID) (c_types.Session, error)
}

type ProfileApi interface {
	IsRegistered(common.Address) (bool, error)
	Register(common.Address) (types.Profile, error)
	GetProfile(common.Address) (types.Profile, error)
}

type PiggyApi interface {
	GetPiggy(address common.Address) (types.Piggy, error)
	GetPiggyFromProfile(address common.Address) (types.Piggy, error)
	GetPiggyFromName(name string) (types.Piggy, error)
	UpdatePiggyName(address common.Address, name string) (types.Piggy, error)
}

type Api interface {
	ProfileApi
	SessionApi
	PiggyApi
}

var (
	_ Api = (*BackendApi)(nil)
)

type BackendApi struct {
	backend *Service
}

func (b *BackendApi) Interface() Api { return b }

func (b *BackendApi) VerifyToken(token string) (*token.Payload, error) {
	payload, err := b.backend.issuer.VerifyToken(token)
	if err != nil {
		return nil, err
	}

	return payload, payload.Valid()
}

func (b *BackendApi) CreateToken(address common.Address) (struct {
	AccessToken    token.Token
	AccessPayload  *token.Payload
	RefreshToken   token.Token
	RefreshPayload *token.Payload
}, error) {
	accessToken, accessPayload, err := b.backend.issuer.IssueAccessToken(address)
	if err != nil {
		return struct {
			AccessToken    token.Token
			AccessPayload  *token.Payload
			RefreshToken   token.Token
			RefreshPayload *token.Payload
		}{}, err
	}
	refreshToken, refreshPayload, err := b.backend.issuer.IssueRefreshToken(address)
	if err != nil {
		return struct {
			AccessToken    token.Token
			AccessPayload  *token.Payload
			RefreshToken   token.Token
			RefreshPayload *token.Payload
		}{}, err
	}

	err = b.backend.session.CreateSession("", refreshToken)
	if err != nil {
		return struct {
			AccessToken    token.Token
			AccessPayload  *token.Payload
			RefreshToken   token.Token
			RefreshPayload *token.Payload
		}{}, err
	}
	return struct {
		AccessToken    token.Token
		AccessPayload  *token.Payload
		RefreshToken   token.Token
		RefreshPayload *token.Payload
	}{
		AccessToken:    accessToken,
		AccessPayload:  accessPayload,
		RefreshToken:   refreshToken,
		RefreshPayload: refreshPayload,
	}, nil
}

func (b *BackendApi) Renew(
	token token.Token,
) (token.Token, *token.Payload, error) {
	payload, err := b.backend.issuer.VerifyToken(string(token))
	if err != nil {
		return "", nil, err
	}

	session, err := b.backend.session.GetSession(payload.ID)
	if err != nil {
		return "", nil, err
	}

	if session.IsBlocked {
		return "", nil, errors.New("session has been blocked")
	}

	if b.backend.clock.Now().After(session.ExpiresAt) {
		return "", nil, errors.New("session has expired")
	}

	return b.backend.issuer.IssueAccessToken(payload.Address)
}

func (b *BackendApi) GetSession(id uuid.UUID) (c_types.Session, error) {
	return c_types.Session{}, errors.ErrUnsupported
}

func (b *BackendApi) Register(
	address common.Address,
) (types.Profile, error) {
	return b.backend.profile.register(address)
}

func (b *BackendApi) GetProfile(
	address common.Address,
) (types.Profile, error) {
	return b.backend.profile.profile(address)
}

func (b *BackendApi) IsRegistered(user common.Address) (bool, error) {
	return b.backend.profile.isRegistered(user)
}

func (b *BackendApi) NotifyNewTransaction(tx common.Hash) error {
	return b.backend.transaction.OnNewTx(tx)
}

func (b *BackendApi) GetPiggy(address common.Address) (types.Piggy, error) {
	return b.backend.piggy.GetPiggy(address)
}
func (b *BackendApi) GetPiggyFromProfile(address common.Address) (types.Piggy, error) {
	return b.backend.piggy.GetPiggyFromProfile(address)
}
func (b *BackendApi) GetPiggyFromName(name string) (types.Piggy, error) {
	return b.backend.piggy.GetPiggyFromName(name)
}
func (b *BackendApi) UpdatePiggyName(address common.Address, name string) (types.Piggy, error) {
	return b.backend.piggy.UpdatePiggyName(address, name)
}
