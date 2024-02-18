package backend

import (
	"context"

	"github.com/Exca-DK/pegism/core/log"
	"github.com/Exca-DK/pegism/core/token"
	"github.com/Exca-DK/pegism/core/types"
	"github.com/Exca-DK/pegism/service/db"
	"github.com/google/uuid"
)

type sessionController struct {
	db     db.Store
	issuer *tokenIssuer
	log    log.Logger
}

func newSessionController(db db.Store, issuer *tokenIssuer) *sessionController {
	return &sessionController{
		db:     db,
		issuer: issuer,
		log:    log.Root().With("module", "session.controller"),
	}
}

func (c *sessionController) CreateSession(ip string, token token.Token) error {
	payload, err := c.issuer.VerifyToken(string(token))
	if err != nil {
		return err
	}
	err = c.db.CreateSession(context.Background(), token, payload)
	if err == nil {
		c.log.Debug(
			"Created new session",
			"address", payload.Address.Hex(),
			"id", payload.ID.String(),
			"exp", payload.ExpiredAtAsTime().String(),
		)
	} else {
		c.log.Warn(
			"Failed creating new session",
			"address", payload.Address.Hex(),
			"id", payload.ID.String(),
			"exp", payload.ExpiredAtAsTime().String(),
			"err", err,
		)
	}
	return err
}

func (c *sessionController) GetSession(id uuid.UUID) (types.Session, error) {
	session, err := c.db.GetSession(context.Background(), id)
	if err != nil {
		return types.Session{}, err
	}
	return session, nil
}
