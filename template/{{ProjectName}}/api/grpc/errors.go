package grpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrGRPCWrongToken               = status.New(codes.Internal, "WrongToken").Err()
	ErrGRPCInvalidArgument          = status.New(codes.InvalidArgument, "InvalidArgument").Err()
	ErrGRPCNotImplemented           = status.New(codes.Unimplemented, "NotImplemented").Err()
	ErrGRPCInternal                 = status.New(codes.Internal, "Internal").Err()
	ErrGRPCUnAuthorized             = status.New(codes.Unauthenticated, "UnAuthorized").Err()
	ErrGRPCResourcePermissionDenied = status.New(codes.PermissionDenied, "Permission denied").Err()
	ErrGRPCUserNotFound             = status.New(codes.NotFound, "UserNotFound").Err()

)
