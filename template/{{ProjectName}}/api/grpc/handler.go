package grpc

import (
	"context"
	api "{{ProjectName}}/api/proto/src"
)

func (Server) RegisterExample(cxt context.Context, req *api.ExampleRequest) (*api.ResponseVoid, error) {
	if req.UserId == 10 {
		return nil, Errors.Internal
	}
	return &api.ResponseVoid{}, nil
}
