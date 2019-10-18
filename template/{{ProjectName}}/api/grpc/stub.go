package grpc

import (
	"google.golang.org/grpc"
	api "{{ProjectName}}/api/proto/src"
	"{{ProjectName}}/internal/service"
	"{{ProjectName}}/pkg"
	"net"
	"google.golang.org/grpc/metadata"
	"context"
	"strconv"
	"time"
	"{{ProjectName}}/pkg/metrics"
	"strings"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/go-grpc-middleware"
)

type Server struct{}

type Option struct {
	Address string
}

func NewGrpcServer(service *service.ExampleServiceImpl, option Option) *Server {
	go listenGrpc(service, option.Address)
	return &Server{}
}

func listenGrpc(service *service.ExampleServiceImpl, address string) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		pkg.Logger.Fatalf("failed to listen: %v", err)
	}
	pkg.Logger.Info("Start listening on address: ", address)
	s := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		rpcInterceptor(),
	)))
	grpc_prometheus.Register(s)
	api.RegisterExampleServer(s, NewHandler(service))
	if err := s.Serve(lis); err != nil {
		pkg.Logger.Fatalf("failed to serve: %v", err)
	}
}

func GetAuthClientId(ctx context.Context) (int32, error) {
	headers, ok := metadata.FromIncomingContext(ctx)
	if ok {
		userIdHeader := headers["user_id"]
		if len(userIdHeader) == 0 {
			return -1, ErrGRPCUnAuthorized
		}
		userID := userIdHeader[0]
		i2, _ := strconv.ParseInt(userID, 10, 64)
		return int32(i2), nil
	}
	return -1, ErrGRPCUnAuthorized
}
func authInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		id, err := GetAuthClientId(ctx)
		if err != nil {
			return nil, err
		}
		ctx = context.WithValue(ctx, "user_id", id)
		return handler(ctx, req)
	}
}

func rpcInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()
		m, err := handler(ctx, req)
		rpcName := extractRpcName(info)
		if err != nil {
			metrics.GetMetrics().MethodErrorCount.WithLabelValues(metrics.ServiceNamespace, rpcName).Add(1)
		} else {
			metrics.GetMetrics().MethodSuccessCount.WithLabelValues(metrics.ServiceNamespace, rpcName).Add(1)
		}
		metrics.GetMetrics().MethodCount.WithLabelValues(metrics.ServiceNamespace, rpcName).Add(1)
		metrics.GetMetrics().MethodDurations.WithLabelValues(metrics.ServiceNamespace, rpcName).
			Observe(float64(time.Since(start).Nanoseconds()))
		return m, togRPCError(err)
	}
}

func extractRpcName(info *grpc.UnaryServerInfo) string {
	return "Request" + strings.Split(info.FullMethod, "/")[2]
}