package grpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrorsDef struct {
	WrongToken      error
	InvalidArgument error
	NotImplemented  error
	Internal        error
}

var Errors = ErrorsDef{
	WrongToken:      status.Errorf(codes.Internal, "WrongToken"),
	InvalidArgument: status.Errorf(codes.InvalidArgument, "InvalidArgument"),
	NotImplemented:  status.Errorf(codes.Unimplemented, "NotImplemented"),
	Internal:        status.Errorf(codes.Internal, "Internal"),
}
