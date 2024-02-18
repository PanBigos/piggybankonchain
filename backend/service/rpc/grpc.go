package rpc

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"net"
	"sync"

	"github.com/Exca-DK/pegism/core/database"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcopentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	grpcRuntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/Exca-DK/pegism/core/clock"
	"github.com/Exca-DK/pegism/core/log"
	"go.opencensus.io/plugin/ocgrpc"
	Grpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type GrpcConfig struct {
	Host           string
	Port           int
	GatewayPort    int
	MaxMsgSizeInMb uint64
	Clock          *clock.Clock
}

func (cfg GrpcConfig) validate() error {
	if cfg.Clock == nil {
		return errors.New("clock has not been provided")
	}

	return nil
}

type GrpcService struct {
	cfg           GrpcConfig
	ctx           context.Context
	cancel        context.CancelFunc
	grpcListener  net.Listener
	grpcServer    *Grpc.Server
	gatewayServer *http.Server

	gatewayListener net.Listener
	gateway         *grpcRuntime.ServeMux

	log log.Logger
	wg  sync.WaitGroup

	connectedRPCClients  map[net.Addr]bool
	clientConnectionLock sync.Mutex

	swagger http.Handler
	header  func(w http.ResponseWriter, s http.Header) error

	clock *clock.Clock
}

func newGrpcService(config GrpcConfig) (*GrpcService, error) {
	if err := config.validate(); err != nil {
		return nil, fmt.Errorf("invalid config provided. %v", err)
	}
	var (
		err     error
		service = &GrpcService{
			cfg:                 config,
			log:                 log.Root().With("module", "grpc"),
			connectedRPCClients: make(map[net.Addr]bool),
			clock:               config.Clock,
		}
	)
	service.ctx, service.cancel = context.WithCancel(context.Background())
	service.grpcListener, err = net.Listen(
		"tcp",
		net.JoinHostPort(service.cfg.Host, strconv.Itoa(service.cfg.Port)),
	)
	if err != nil {
		return nil, err
	}
	if service.cfg.GatewayPort != 0 {
		service.gatewayListener, err = net.Listen(
			"tcp",
			net.JoinHostPort(service.cfg.Host, strconv.Itoa(service.cfg.GatewayPort)),
		)
		if err != nil {
			return nil, err
		}
	}

	if service.cfg.MaxMsgSizeInMb == 0 {
		service.log.Debug("Increasing max frame size to 8 mb ")
		service.cfg.MaxMsgSizeInMb = 1024 * 1024 * 8
	} else {
		service.cfg.MaxMsgSizeInMb *= 1024 * 1024
	}
	opts := []Grpc.ServerOption{
		Grpc.StatsHandler(&ocgrpc.ServerHandler{}),
		Grpc.StreamInterceptor(middleware.ChainStreamServer(
			recovery.StreamServerInterceptor(),
			grpcprometheus.StreamServerInterceptor,
			grpcopentracing.StreamServerInterceptor(),
			service.validatorStreamConnectionInterceptor,
		)),
		Grpc.UnaryInterceptor(middleware.ChainUnaryServer(
			recovery.UnaryServerInterceptor(),
			grpcprometheus.UnaryServerInterceptor,
			grpcopentracing.UnaryServerInterceptor(),
			service.clientUnaryClientInfoInterceptor,
			service.clientUnaryErrorInterceptor,
		)),
		Grpc.MaxRecvMsgSize(int(service.cfg.MaxMsgSizeInMb)),
	}

	service.grpcServer = Grpc.NewServer(opts...)
	service.gateway = grpcRuntime.NewServeMux()
	return service, nil
}

func (s *GrpcService) RegisterRpcService(f func(srv *Grpc.Server)) { f(s.grpcServer) }

func (s *GrpcService) RegisterGatewayService(
	f func(ctx context.Context, mux *grpcRuntime.ServeMux, endpoint string) error,
) error {
	return f(s.ctx, s.gateway, s.grpcListener.Addr().String())
}

func (s *GrpcService) WithSwagger(h http.Handler) *GrpcService {
	s.swagger = h
	return s
}

func (s *GrpcService) WithHeader(h func(http.ResponseWriter, http.Header) error) *GrpcService {
	s.header = h
	return s
}

func (s *GrpcService) Start() {
	// Register reflection service on gRPC server.
	reflection.Register(s.grpcServer)

	if s.grpcListener != nil {
		s.wg.Add(1)
		go s.startGrpc()
	}

	if s.gatewayListener != nil {
		s.wg.Add(1)
		go s.startGateway()
	}
}

func (s *GrpcService) startGrpc() {
	defer s.wg.Done()
	if s.grpcListener != nil {
		s.log.Info(
			"grpc listening",
			"endpoint", s.grpcListener.Addr().String(),
		)
		if err := s.grpcServer.Serve(s.grpcListener); err != nil {
			s.log.Error("grpc serve failure", "err", err)
		}
	}
}

func (s *GrpcService) startGateway() {
	defer s.wg.Done()
	swagger := s.swagger
	s.gatewayServer = &http.Server{
		Addr: s.gatewayListener.Addr().String(),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if s.header != nil {
				if err := s.header(w, r.Header); err != nil {
					s.log.Info(
						"Ignored http request",
						"err", err,
						"addr", r.RemoteAddr,
					)
					return
				}
			}

			if r.Method == http.MethodOptions {
				return
			}

			if strings.HasPrefix(r.URL.Path, "/docs") {
				if swagger != nil {
					swagger.ServeHTTP(w, r)
				}
			} else {
				s.gateway.ServeHTTP(w, r)
			}
		}),
	}

	s.log.Info(
		"Gateway server listening",
		"endpoint", s.gatewayListener.Addr().String(),
	)
	if err := s.gatewayServer.Serve(s.gatewayListener); err != nil {
		s.log.Error("mux serve failure", "err", err)
	}
}

// Tries to stop the server gracefully
func (s *GrpcService) Stop() error {
	s.cancel()
	if s.grpcListener != nil {
		s.log.Debug("Initiated graceful stop of gRPC server")
		s.grpcServer.GracefulStop()
	}
	if s.gatewayListener != nil {
		s.log.Debug("Initiated graceful stop of gRPC Gateway server")
		s.gatewayServer.Shutdown(context.Background())
	}
	s.wg.Wait()
	if s.grpcListener != nil {
		s.grpcListener.Close()
	}
	if s.gatewayListener != nil {
		s.gatewayListener.Close()
	}
	return nil
}

func (s *GrpcService) validatorStreamConnectionInterceptor(
	srv interface{},
	ss Grpc.ServerStream,
	_ *Grpc.StreamServerInfo,
	handler Grpc.StreamHandler,
) error {
	s.logNewClientConnection(ss.Context())
	return handler(srv, ss)
}

func (s *GrpcService) validatorUnaryConnectionInterceptor(
	ctx context.Context,
	req interface{},
	_ *Grpc.UnaryServerInfo,
	handler Grpc.UnaryHandler,
) (interface{}, error) {
	s.logNewClientConnection(ctx)
	return handler(ctx, req)
}

func (s *GrpcService) logNewClientConnection(ctx context.Context) {
	if clientInfo, ok := peer.FromContext(ctx); ok {
		// Check if we have not yet observed this grpc client connection
		s.clientConnectionLock.Lock()
		defer s.clientConnectionLock.Unlock()
		if !s.connectedRPCClients[clientInfo.Addr] {
			s.log.Debug(
				"gRPC client connected",
				"addr", clientInfo.Addr.String(),
			)
			s.connectedRPCClients[clientInfo.Addr] = true
		}
	}
}

func (s *GrpcService) clientUnaryErrorInterceptor(
	ctx context.Context,
	req interface{},
	info *Grpc.UnaryServerInfo,
	handler Grpc.UnaryHandler,
) (interface{}, error) {
	result, err := handler(ctx, req)
	if err != nil {
		_, ok := status.FromError(err)
		if ok {
			return nil, err
		}
		converted, ok := statusFromKnown(err)
		if ok {
			return nil, converted.Err()
		}
		return nil, status.Error(codes.Unknown, err.Error())
	}
	return result, nil
}

func (s *GrpcService) clientUnaryClientInfoInterceptor(
	ctx context.Context,
	req interface{},
	info *Grpc.UnaryServerInfo,
	handler Grpc.UnaryHandler,
) (interface{}, error) {
	start := s.clock.Now()
	metadata := extractGrpcMetadata(ctx)
	result, err := handler(ctx, req)
	s.log.Trace(
		"Invoked RPC method",
		"userAgent", metadata.UserAgent,
		"clientIp", metadata.ClientIP,
		"method", info.FullMethod,
		"duration", s.clock.Since(start).String(),
		"err", err,
	)
	return result, err
}

const (
	grpcGatewayUserAgentHeader = "grpcgateway-user-agent"
	userAgentHeader            = "user-agent"
	xForwardedForHeader        = "x-forwarded-for"
)

type grpcMetadata struct {
	UserAgent string
	ClientIP  string
}

func extractGrpcMetadata(ctx context.Context) *grpcMetadata {
	mtdt := &grpcMetadata{}

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if userAgents := md.Get(grpcGatewayUserAgentHeader); len(userAgents) > 0 {
			mtdt.UserAgent = userAgents[0]
		}

		if userAgents := md.Get(userAgentHeader); len(userAgents) > 0 {
			mtdt.UserAgent = userAgents[0]
		}

		if clientIPs := md.Get(xForwardedForHeader); len(clientIPs) > 0 {
			mtdt.ClientIP = clientIPs[0]
		}
	}

	if p, ok := peer.FromContext(ctx); ok {
		mtdt.ClientIP = p.Addr.String()
	}

	return mtdt
}

type StatusCode uint

var (
	UnknownDbCode uint = 1000
)

var (
	ErrNotFound  = status.New(codes.NotFound, codes.NotFound.String())
	ErrDuplicate = status.New(codes.AlreadyExists, codes.AlreadyExists.String())

	ErrDbUnknown = status.New(codes.Code(UnknownDbCode), "unknown db error")
)

func statusFromKnown(err error) (*status.Status, bool) {
	var (
		status *status.Status
		ok     bool
	)
	status, ok = tryStatusFromDb(err)
	if ok {
		return status, true
	}

	return nil, false
}

func tryStatusFromDb(err error) (*status.Status, bool) {
	wrapped, ok := err.(database.DatabaseError)
	if !ok {
		return nil, false
	}

	switch wrapped {
	case database.ErrDbNotFound:
		return ErrNotFound, true
	case database.ErrDbNotUnique:
		return ErrDuplicate, true
	}

	return ErrDbUnknown, true
}
