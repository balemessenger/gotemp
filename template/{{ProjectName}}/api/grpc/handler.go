package grpc

import (
	"context"
	api "{{ProjectName}}/api/proto/src"
	"{{ProjectName}}/internal/service"
)

type Handler struct {
	service *service.ExampleServiceImpl
}

func NewHandler(service *service.ExampleServiceImpl) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RpcExample(ctx context.Context, req *api.ExampleRequest) (*api.ResponseVoid, error) {
	return h.service.RpcExample(ctx, req)
}
