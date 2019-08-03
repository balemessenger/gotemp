package grpc

import (
	"context"
	api "{{ProjectName}}/api/proto/src"
	"{{ProjectName}}/pkg"
)

type Handler struct {
	log *pkg.Logger
}

func NewHandler(log *pkg.Logger) *Handler {
	return &Handler{log: log}
}

func (h *Handler) RegisterExample(cxt context.Context, req *api.ExampleRequest) (*api.ResponseVoid, error) {
	if req.UserId == 10 {
		return nil, Errors.Internal
	}
	return &api.ResponseVoid{}, nil
}
