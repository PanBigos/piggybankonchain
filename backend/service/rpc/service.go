package rpc

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/Exca-DK/pegism/service/backend"
	"github.com/Exca-DK/pegism/service/rpc/auth"
	"github.com/Exca-DK/pegism/service/rpc/piggy"
	"github.com/Exca-DK/pegism/service/rpc/profile"

	"github.com/Exca-DK/pegism/core/types"

	protoServices "github.com/Exca-DK/pegism/gen/go/proto/service"
	grpcRuntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	googleGrpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Config struct {
	GrpcConfig GrpcConfig

	Api backend.Api
}

type RpcService struct {
	cfg           Config
	authServer    *auth.AuthServer
	profileServer *profile.ProfileServer
	piggyServer   *piggy.PiggyServer

	grpc *GrpcService
}

func NewService(config *Config) (*RpcService, error) {
	service := &RpcService{cfg: *config}
	grpcService, err := newGrpcService(config.GrpcConfig)
	if err != nil {
		return nil, err
	}
	service.grpc = grpcService
	service.grpc = service.grpc.WithHeader(func(w http.ResponseWriter, h http.Header) error {
		if !service.allowedOrigin(h.Get("Origin")) {
			return errors.New("unallowed origin")
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		return nil
	})

	if err := service.setupAuth(); err != nil {
		return nil, fmt.Errorf("auth server setup failre. Error: %v", err)
	}
	if err := service.setupProfile(); err != nil {
		return nil, fmt.Errorf("profile server setup failre. Error: %v", err)
	}
	if err := service.setupPiggy(); err != nil {
		return nil, fmt.Errorf("piggy server setup failre. Error: %v", err)
	}

	return service, nil
}

func (service *RpcService) GetAuthServer() *auth.AuthServer {
	return service.authServer
}

func (service *RpcService) GetProfileServer() *profile.ProfileServer {
	return service.profileServer
}

func (service *RpcService) GetPiggyServer() *profile.ProfileServer {
	return service.profileServer
}

func (service *RpcService) setupProfile() error {
	creds := googleGrpc.WithTransportCredentials(insecure.NewCredentials())
	service.profileServer = profile.NewServer(service.cfg.Api)
	service.grpc.RegisterRpcService(func(srv *googleGrpc.Server) {
		protoServices.RegisterProfileServiceServer(srv, service.profileServer)
	})

	if err := service.grpc.RegisterGatewayService(
		func(ctx context.Context, mux *grpcRuntime.ServeMux, endpoint string) error {
			return protoServices.RegisterProfileServiceHandlerFromEndpoint(
				ctx,
				mux,
				endpoint,
				[]googleGrpc.DialOption{creds},
			)
		},
	); err != nil {
		return fmt.Errorf("gateway registration failure. error: %w", err)
	}
	return nil
}

func (service *RpcService) setupAuth() error {
	creds := googleGrpc.WithTransportCredentials(insecure.NewCredentials())
	service.authServer = auth.NewServer(
		service.cfg.Api,
		types.RegistrationMessageFactory{}.GenerateNewTemplate("pegism.auth", "registration"),
		service.cfg.GrpcConfig.Clock,
	)
	service.grpc.RegisterRpcService(func(srv *googleGrpc.Server) {
		protoServices.RegisterAuthorizationServiceServer(srv, service.authServer)
	})

	if err := service.grpc.RegisterGatewayService(
		func(ctx context.Context, mux *grpcRuntime.ServeMux, endpoint string) error {
			return protoServices.RegisterAuthorizationServiceHandlerFromEndpoint(
				ctx,
				mux,
				endpoint,
				[]googleGrpc.DialOption{creds},
			)
		},
	); err != nil {
		return fmt.Errorf("gateway registration failure. error: %w", err)
	}
	return nil
}

func (service *RpcService) setupPiggy() error {
	creds := googleGrpc.WithTransportCredentials(insecure.NewCredentials())
	service.piggyServer = piggy.NewServer(service.cfg.Api)
	service.grpc.RegisterRpcService(func(srv *googleGrpc.Server) {
		protoServices.RegisterPiggyServiceServer(srv, service.piggyServer)
	})

	if err := service.grpc.RegisterGatewayService(
		func(ctx context.Context, mux *grpcRuntime.ServeMux, endpoint string) error {
			return protoServices.RegisterPiggyServiceHandlerFromEndpoint(
				ctx,
				mux,
				endpoint,
				[]googleGrpc.DialOption{creds},
			)
		},
	); err != nil {
		return fmt.Errorf("gateway registration failure. error: %w", err)
	}
	return nil
}

// TODO implement
func (s *RpcService) Status() error { return nil }

func (s *RpcService) Start() error {
	s.grpc.Start()
	return nil
}

// Tries to stop the server gracefully
func (s *RpcService) Stop() error {
	return s.grpc.Stop()
}

func (s *RpcService) allowedOrigin(origin string) bool {
	return true // TODO fix later
	// if s.cfg.AllowedOrigin == "*" {
	// 	return true
	// }
	// if matched, _ := regexp.MatchString(s.cfg.AllowedOrigin, origin); matched {
	// 	return true
	// }
	// return false
}
