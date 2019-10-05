package grpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrorsDef struct {
	InternalError error
}

var Errors = ErrorsDef{
	InternalError: status.Errorf(codes.Internal, "Internal error"),
}
