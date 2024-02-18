package piggy

import (
	"context"

	"github.com/Exca-DK/pegism/service/backend"
	"github.com/Exca-DK/pegism/service/rpc/auth"

	"github.com/Exca-DK/pegism/core/token"

	"github.com/Exca-DK/pegism/core/log"
	"github.com/ethereum/go-ethereum/common"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Exca-DK/pegism/gen/go/proto/service"
	v1 "github.com/Exca-DK/pegism/gen/go/proto/v1"
)

var (
	_     service.PiggyServiceServer = (*PiggyServer)(nil)
	empty                            = &emptypb.Empty{}
)

// TODO trace every request

type backendApi interface {
	backend.PiggyApi
	VerifyToken(string) (*token.Payload, error)
}

type PiggyServer struct {
	service.UnimplementedPiggyServiceServer
	log log.Logger
	api backendApi
}

func NewServer(api backendApi) *PiggyServer {
	server := &PiggyServer{
		log: log.Root(),
		api: api,
	}
	return server
}

func (srv *PiggyServer) GetPiggy(
	ctx context.Context,
	req *v1.GetPiggyRequest,
) (*v1.GetPiggyResponse, error) {
	address := common.HexToAddress(req.Address)
	piggy, err := srv.api.GetPiggy(address)
	if err != nil {
		srv.log.Warn(
			"Unexpected error at ProfileServer.IsUserRegistered",
			"address", address.Hex(),
			"err", err,
		)
		return nil, err
	}
	protoPiggy, err := piggy.ToProto()
	if err != nil {
		return nil, err
	}
	return &v1.GetPiggyResponse{Piggy: protoPiggy}, nil
}

func (srv *PiggyServer) GetPiggyFromProfile(
	ctx context.Context,
	req *v1.GetPiggyRequest,
) (*v1.GetPiggyResponse, error) {
	address := common.HexToAddress(req.Address)
	piggy, err := srv.api.GetPiggyFromProfile(address)
	if err != nil {
		srv.log.Warn(
			"Unexpected error",
			"address", address.Hex(),
			"err", err,
		)
		return nil, err
	}
	protoPiggy, err := piggy.ToProto()
	if err != nil {
		return nil, err
	}
	return &v1.GetPiggyResponse{Piggy: protoPiggy}, nil
}

func (srv *PiggyServer) GetPiggyFromName(
	ctx context.Context,
	req *v1.GetPiggyFromNameRequest,
) (*v1.GetPiggyFromNameResponse, error) {
	piggy, err := srv.api.GetPiggyFromName(req.Name)
	if err != nil {
		srv.log.Warn(
			"Unexpected error",
			"name", req.Name,
			"err", err,
		)
		return nil, err
	}
	protoPiggy, err := piggy.ToProto()
	if err != nil {
		return nil, err
	}
	return &v1.GetPiggyFromNameResponse{Piggy: protoPiggy}, nil
}

func (srv *PiggyServer) UpdatePiggyName(
	ctx context.Context,
	req *v1.UpdatePiggyNameRequest,
) (*v1.UpdatePiggyNameResponse, error) {
	address := common.HexToAddress(req.Address)
	token, err := srv.recoverToken(ctx)
	if err != nil {
		return nil, err
	}
	if token.Address != address {
		return nil, status.New(codes.PermissionDenied, "request address doesn't match identity").
			Err()
	}
	piggy, err := srv.api.UpdatePiggyName(address, req.Name)
	if err != nil {
		srv.log.Warn(
			"Unexpected error at ProfileServer.IsUserRegistered",
			"address", address.Hex(),
			"err", err,
		)
		return nil, err
	}
	protoPiggy, err := piggy.ToProto()
	if err != nil {
		return nil, err
	}
	return &v1.UpdatePiggyNameResponse{Piggy: protoPiggy}, nil
}

func (srv *PiggyServer) recoverToken(ctx context.Context) (*token.Payload, error) {
	accessTokenRaw, err := auth.GrpcRecoverAuthTokenHeader(ctx)
	if err != nil {
		return nil, auth.UnauthenticatedError(err)
	}
	accessToken, err := srv.api.VerifyToken(accessTokenRaw)
	if err != nil {
		return nil, auth.UnauthenticatedError(err)
	}
	return accessToken, nil
}
