package auth

import (
	"context"
	"time"

	"github.com/Exca-DK/pegism/core/clock"
	"github.com/Exca-DK/pegism/core/token"
	ctypes "github.com/Exca-DK/pegism/core/types"
	"github.com/Exca-DK/pegism/service/backend"

	"github.com/Exca-DK/pegism/core/log"
	"github.com/ethereum/go-ethereum/common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Exca-DK/pegism/gen/go/proto/service"
	v1 "github.com/Exca-DK/pegism/gen/go/proto/v1"
)

type AuthServer struct {
	service.UnimplementedAuthorizationServiceServer
	log log.Logger
	api backend.SessionApi

	authTemplate *ctypes.RegistrationMessageTemplate
	clock        *clock.Clock
}

func NewServer(
	api backend.SessionApi,
	template *ctypes.RegistrationMessageTemplate,
	clock *clock.Clock,
) *AuthServer {
	return &AuthServer{
		UnimplementedAuthorizationServiceServer: service.UnimplementedAuthorizationServiceServer{},
		log:                                     log.Root(),
		api:                                     api,
		authTemplate:                            template,
		clock:                                   clock,
	}
}

func (srv *AuthServer) Authorize(
	ctx context.Context,
	req *v1.AuthRequest,
) (*v1.AuthResponse, error) {
	var (
		err error
	)

	_, header, ok := srv.authTemplate.Recover(ctypes.AuthMessage(req.GetArgs().GetMsg()))
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "failed recovering signed header")
	}

	err = ValidateAuth(req.GetArgs(), header)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, err.Error())
	}

	addr := header.Acquirer

	token, err := srv.api.CreateToken(addr)
	if err != nil {
		srv.log.Debug("failed token creation",
			"address", addr.Hex(),
			"err", err,
		)
		return nil, status.Error(codes.Internal, "token creation failed")
	}

	return &v1.AuthResponse{
		Data: &v1.AuthResponse_AuthorizationDataContainer{
			AccessToken:           string(token.AccessToken),
			RefreshToken:          string(token.RefreshToken),
			AccessTokenExpiresAt:  timestamppb.New(token.AccessPayload.ExpiredAtAsTime()),
			RefreshTokenExpiresAt: timestamppb.New(token.RefreshPayload.ExpiredAtAsTime()),
		},
	}, nil
}

func (srv *AuthServer) GetAuthMessage(
	ctx context.Context,
	req *v1.AuthMessageRequest,
) (*v1.AuthMessageResponse, error) {
	deadline := srv.clock.Now().Add(30 * time.Second)
	msg := srv.authTemplate.NewMessage(common.HexToAddress(req.GetAddress()), deadline)
	return &v1.AuthMessageResponse{
		Message: &v1.UnsignedAuthMessage{
			Content:  string(msg),
			Deadline: timestamppb.New(deadline),
		},
	}, nil
}

func (srv *AuthServer) Refresh(
	ctx context.Context,
	req *v1.RefreshMessageRequest,
) (*v1.RefreshMessageResponse, error) {
	token, payload, err := srv.api.Renew(token.Token(req.RefreshToken))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &v1.RefreshMessageResponse{
		AccessToken: string(token),
		ExpiresAt:   timestamppb.New(payload.ExpiredAtAsTime()),
	}, nil
}
