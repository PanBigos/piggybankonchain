package backend

import (
	"context"

	"github.com/Exca-DK/pegism/core/log"

	"github.com/Exca-DK/pegism/service/db"
	"github.com/Exca-DK/pegism/service/types"
	"github.com/ethereum/go-ethereum/common"
)

type profileController struct {
	store  db.Store
	logger log.Logger
}

func newProfile(store db.Store) *profileController {
	return &profileController{
		store:  store,
		logger: log.Root(),
	}
}

func (p *profileController) isRegistered(address common.Address) (bool, error) {
	return p.store.Registered(context.Background(), address)
}

func (p *profileController) register(
	address common.Address,
) (types.Profile, error) {
	return p.store.CreateProfile(context.Background(), address)
}

func (p *profileController) profile(address common.Address) (types.Profile, error) {
	return p.store.GetProfile(context.Background(), address)
}
