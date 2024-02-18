package backend

import (
	"context"

	"github.com/Exca-DK/pegism/service/db"
	"github.com/Exca-DK/pegism/service/types"
	"github.com/ethereum/go-ethereum/common"
)

type piggyController struct {
	store db.Store
}

func newPiggyController(store db.Store) *piggyController {
	return &piggyController{
		store: store,
	}
}

func (p *piggyController) AddPiggy(piggy types.Piggy) (types.Piggy, error) {
	return p.store.CreatePiggy(context.Background(), piggy)
}

func (p *piggyController) AddPiggyMessage(msg types.Message) (types.Message, error) {
	return p.store.AddMessage(context.Background(), msg)
}

func (p *piggyController) UpdatePiggyName(
	address common.Address,
	name string,
) (types.Piggy, error) {
	return p.store.UpdatePiggyName(context.Background(), address, name)
}

func (p *piggyController) GetPiggy(address common.Address) (types.Piggy, error) {
	return p.store.GetPiggy(context.Background(), address)
}

func (p *piggyController) GetPiggyFromProfile(address common.Address) (types.Piggy, error) {
	return p.store.GetPiggyFromProfile(context.Background(), address)
}

func (p *piggyController) GetPiggyFromName(name string) (types.Piggy, error) {
	return p.store.GetPiggyFromName(context.Background(), name)
}
