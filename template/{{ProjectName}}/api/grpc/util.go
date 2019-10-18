package grpc

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	_ "github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"{{ProjectName}}/internal/repositories"
	"{{ProjectName}}/pkg"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

var toGRPCErrorMap = map[error]error{
	repositories.ErrUserNotFound:    ErrGRPCUserNotFound,
}

func getStackTrace(err error) string {
	if err, ok := err.(stackTracer); ok {
		return fmt.Sprintf("%+v", err.StackTrace())
	}
	return ""
}

func togRPCErrorM(err error) error {
	// let gRPC server convert to codes.Canceled, codes.DeadlineExceeded
	if err == context.Canceled || err == context.DeadlineExceeded {
		return err
	}
	grpcErr, ok := toGRPCErrorMap[err]
	if !ok {
		pkg.Logger.WithField("stacktrace", getStackTrace(err)).Error(err)
		return status.Error(codes.Unknown, err.Error())
	}
	return grpcErr
}

func togRPCError(err error) error {
	switch err := err.(type) {
	case nil:
		return nil
	//case cassandra.ErrMessageNotFound:
	//	return ErrGRPCMessageNotFound

	default:
		return togRPCErrorM(err)
	}
}
