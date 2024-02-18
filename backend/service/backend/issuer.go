package backend

import (
	"time"

	"github.com/Exca-DK/pegism/core/clock"
	"github.com/Exca-DK/pegism/core/token"
	"github.com/ethereum/go-ethereum/common"
)

type tokenIssuer struct {
	factory         token.TokenFactory
	acessDuration   time.Duration
	refreshDuration time.Duration
}

func newTokenController(
	secret string,
	acessDuration time.Duration,
	refreshDuration time.Duration,
	clock *clock.Clock,
) (*tokenIssuer, error) {
	factory, err := token.NewJwtFactory(secret, clock)
	if err != nil {
		return nil, err
	}
	return &tokenIssuer{
		factory:         factory,
		acessDuration:   acessDuration,
		refreshDuration: refreshDuration,
	}, nil
}

func (controller *tokenIssuer) VerifyToken(token string) (*token.Payload, error) {
	return controller.factory.VerifyToken(token)
}

func (controller *tokenIssuer) IssueAccessToken(
	username common.Address,
) (token.Token, *token.Payload, error) {
	return controller.factory.CreateToken(username.Hex(), controller.getAccessIssueDuration())
}

func (controller *tokenIssuer) IssueRefreshToken(
	username common.Address,
) (token.Token, *token.Payload, error) {
	return controller.factory.CreateToken(username.Hex(), controller.getRefreshIssueDuration())
}

func (controller *tokenIssuer) getAccessIssueDuration() time.Duration {
	return controller.acessDuration
}

func (controller *tokenIssuer) getRefreshIssueDuration() time.Duration {
	return controller.refreshDuration
}
