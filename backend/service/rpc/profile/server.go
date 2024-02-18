package profile

import (
	"context"

	"github.com/Exca-DK/pegism/service/backend"
	"github.com/Exca-DK/pegism/service/rpc/auth"
	"github.com/Exca-DK/pegism/service/types"

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
	_     service.ProfileServiceServer = (*ProfileServer)(nil)
	empty                              = &emptypb.Empty{}
)

// TODO trace every request

type backendApi interface {
	backend.ProfileApi
	VerifyToken(string) (*token.Payload, error)
}

type ProfileServer struct {
	service.UnimplementedProfileServiceServer
	log log.Logger
	api backendApi
}

func NewServer(api backendApi) *ProfileServer {
	server := &ProfileServer{
		log: log.Root(),
		api: api,
	}
	return server
}

func (srv *ProfileServer) IsRegistered(
	ctx context.Context,
	req *v1.IsRegisteredRequest,
) (*v1.IsRegisteredResponse, error) {
	address := common.HexToAddress(req.Address)
	registered, err := srv.isRegistered(address)
	if err != nil {
		return nil, err
	}
	return &v1.IsRegisteredResponse{
		Registered: registered,
	}, nil
}

func (srv *ProfileServer) isRegistered(address common.Address) (bool, error) {
	v, err := srv.api.IsRegistered(address)
	if err != nil {
		srv.log.Warn(
			"Unexpected error at ProfileServer.IsUserRegistered",
			"address", address.Hex(),
			"err", err,
		)
		return false, status.Error(codes.Internal, "something went wrong")
	}
	return v, nil
}

func (srv *ProfileServer) Register(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	token, err := srv.recoverToken(ctx)
	if err != nil {
		return nil, err
	}
	address := token.Address

	registered, err := srv.isRegistered(address)
	if err != nil {
		return nil, err
	}

	if registered {
		return nil, status.Error(codes.AlreadyExists, "already registered")
	}

	if _, err := srv.api.Register(address); err != nil {
		srv.log.Warn(
			"Failed streamer registration",
			"address", address.Hex(),
			"err", err,
		)
		return nil, status.Error(codes.Aborted, "something went wrong")
	}

	srv.log.Info(
		"Registered new profile",
		"address", address.Hex(),
	)

	return empty, nil
}

func (srv *ProfileServer) GetProfile(
	ctx context.Context,
	req *v1.GetProfileRequest,
) (*v1.GetProfileResponse, error) {
	address := common.HexToAddress(req.Address)

	profile, status := srv.getProfile(address)
	if status != nil {
		return nil, status.Err()
	}
	adapter := profileAdapter{}
	adapted, err := adapter.ToProto(profile)
	if err != nil {
		return nil, err
	}
	return &v1.GetProfileResponse{Profile: adapted}, nil
}

func (srv *ProfileServer) getProfile(
	address common.Address,
) (types.Profile, *status.Status) {
	registered, err := srv.api.IsRegistered(address)
	if err != nil {
		return types.Profile{}, status.New(
			codes.Internal,
			"something went wrong: code(100)",
		)
	}
	if !registered {
		return types.Profile{}, status.New(codes.NotFound, "not registered")
	}

	profile, err := srv.api.GetProfile(address)
	if err != nil {
		srv.log.Warn(
			"Failed getting streamer profile",
			"address", address.Hex(),
			"err", err,
		)
		return types.Profile{}, status.New(
			codes.Internal,
			"something went wrong: code(102)",
		)
	}

	return profile, nil
}

func (srv *ProfileServer) authorizeUser(ctx context.Context, address common.Address) error {
	accessToken, err := srv.recoverToken(ctx)
	if err != nil {
		return err
	}
	if accessToken.Address != address {
		srv.log.Warn(
			"Request with invalid identity",
			"address", address.Hex(),
			"identity", accessToken.Address.Hex(),
		)
		return status.Error(codes.Unauthenticated, "invalid credentials")
	}
	return nil
}

func (srv *ProfileServer) recoverToken(ctx context.Context) (*token.Payload, error) {
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
