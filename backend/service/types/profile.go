package types

import (
	"math/big"
	"time"

	"github.com/Exca-DK/pegism/core/option"
	v1 "github.com/Exca-DK/pegism/gen/go/proto/v1"
	"github.com/ethereum/go-ethereum/common"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Profile struct {
	Address   common.Address
	Piggies   []Piggy
	CreatedAt time.Time
}

type Message struct {
	Hash    common.Hash
	Address common.Address
	Token   common.Address
	Amount  *big.Int
	Fee     *big.Int
	Nick    string
	Content string
	AddedAt time.Time
}

func (p *Message) ToProto() *v1.PiggyMessage {
	return &v1.PiggyMessage{
		Address: p.Address.Hex(),
		Token:   p.Token.Hex(),
		Amount:  p.Amount.String(),
		Fee:     p.Fee.String(),
		Content: p.Content,
		Nick:    p.Nick,
	}
}

type Piggy struct {
	Address        common.Address
	FromAddress    common.Address
	ProfileAddress common.Address
	CreatedAt      time.Time
	AddedAt        time.Time
	UnlocksAt      time.Time
	Name           option.Option[string]
	Messages       []Message
}

func (p *Piggy) ToProto() (*v1.Piggy, error) {
	var name *string
	if p.Name.Some() {
		v := p.Name.MustUnwrap()
		name = &v
	}
	msgs := make([]*v1.PiggyMessage, len(p.Messages))
	for i, msg := range p.Messages {
		msgs[i] = msg.ToProto()
	}
	return &v1.Piggy{
		Address:     p.Address.Hex(),
		Creator:     p.FromAddress.Hex(),
		Owner:       p.ProfileAddress.Hex(),
		CreatedAt:   timestamppb.New(p.CreatedAt),
		AddedAt:     timestamppb.New(p.AddedAt),
		UnlocksAt:   timestamppb.New(p.UnlocksAt),
		DisplayName: name,
		Messages:    msgs,
	}, nil
}
