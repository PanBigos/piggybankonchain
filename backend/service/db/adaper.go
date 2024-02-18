package db

import (
	"context"
	"math/big"

	"github.com/Exca-DK/pegism/core/option"
	"github.com/Exca-DK/pegism/service/db/sqlc"
	"github.com/Exca-DK/pegism/service/types"
	"github.com/ethereum/go-ethereum/common"
)

type messageAdapter struct{}

func (p *messageAdapter) ToModel(piggy types.Message) sqlc.Message {
	return sqlc.Message{
		TransactionHash: piggy.Hash.Hex(),
		Address:         piggy.Address.Hex(),
		Token:           piggy.Token.String(),
		Amount:          piggy.Amount.String(),
		Fee:             piggy.Fee.String(),
		Content:         piggy.Content,
		Nick:            piggy.Nick,
		AddedAt:         piggy.AddedAt,
	}
}

func (p *messageAdapter) FromModel(piggy sqlc.Message) types.Message {
	amount, _ := big.NewInt(0).SetString(piggy.Amount, 10)
	fee, _ := big.NewInt(0).SetString(piggy.Fee, 10)
	return types.Message{
		Address: common.HexToAddress(piggy.Address),
		Token:   common.HexToAddress(piggy.Token),
		Amount:  amount,
		Fee:     fee,
		Nick:    piggy.Nick,
		Content: piggy.Content,
	}
}

type piggyAdapter struct{}

func (p *piggyAdapter) ToModel(piggy types.Piggy) sqlc.Piggy {
	model := sqlc.Piggy{
		Address:        piggy.Address.Hex(),
		FromAddress:    piggy.FromAddress.Hex(),
		ProfileAddress: piggy.ProfileAddress.Hex(),
		CreatedAt:      piggy.CreatedAt,
		AddedAt:        piggy.AddedAt,
		UnlocksAt:      piggy.UnlocksAt,
	}
	if piggy.Name.Some() {
		model.Name.String = piggy.Name.Value()
		model.Name.Valid = true
	}
	return model
}

func (p *piggyAdapter) FromModel(piggy sqlc.Piggy, piggy_msgs []sqlc.Message) types.Piggy {
	var (
		name option.Option[string]
	)
	if piggy.Name.Valid {
		name = option.Some(piggy.Name.String)
	}
	adapter := messageAdapter{}
	msgs := make([]types.Message, len(piggy_msgs))
	for i, msg := range piggy_msgs {
		msgs[i] = adapter.FromModel(msg)
	}
	return types.Piggy{
		Address:        common.HexToAddress(piggy.Address),
		FromAddress:    common.HexToAddress(piggy.FromAddress),
		ProfileAddress: common.HexToAddress(piggy.ProfileAddress),
		CreatedAt:      piggy.CreatedAt,
		AddedAt:        piggy.AddedAt,
		UnlocksAt:      piggy.UnlocksAt,
		Name:           name,
		Messages:       msgs,
	}
}

type profileAdapter struct {
	q sqlc.Querier
}

func (p *profileAdapter) ToModel(profile types.Profile) sqlc.Profile {
	return sqlc.Profile{
		Address:   profile.Address.Hex(),
		CreatedAt: profile.CreatedAt,
	}
}

func (p *profileAdapter) FromModel(model sqlc.Profile) (types.Profile, error) {
	profile := types.Profile{
		Address:   common.HexToAddress(model.Address),
		CreatedAt: model.CreatedAt,
	}
	piggies, err := p.q.GetPiggies(context.Background(), profile.Address.Hex())
	if err != nil {
		return types.Profile{}, err
	}
	profile.Piggies = make([]types.Piggy, len(piggies))
	for i, sqlc_piggy := range piggies {
		msgs, err := p.q.GetMessagesByAddress(context.Background(), sqlc_piggy.Address)
		if err != nil {
			return types.Profile{}, err
		}
		adapter := piggyAdapter{}
		piggy := adapter.FromModel(sqlc_piggy, msgs)
		profile.Piggies[i] = piggy
	}

	return profile, nil
}
