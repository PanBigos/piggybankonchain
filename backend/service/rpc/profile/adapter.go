package profile

import (
	v1 "github.com/Exca-DK/pegism/gen/go/proto/v1"
	"github.com/Exca-DK/pegism/service/types"
)

type profileAdapter struct {
}

func (profileAdapter) ToProto(profile types.Profile) (*v1.Profile, error) {
	piggies := make([]*v1.Piggy, len(profile.Piggies))
	for i, piggy := range profile.Piggies {
		protoPiggy, err := piggy.ToProto()
		if err != nil {
			return nil, err
		}
		piggies[i] = protoPiggy
	}
	return &v1.Profile{
		Address: profile.Address.Hex(),
		Piggies: piggies,
	}, nil
}
