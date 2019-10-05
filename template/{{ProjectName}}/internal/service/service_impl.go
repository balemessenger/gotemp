package service

import (
	"context"
	api "{{ProjectName}}/api/proto/src"
)

type ExampleServiceImpl struct {
	// Inject dependencies
	option Option
}

type Option struct {
}

func NewExampleService() *ExampleServiceImpl {
	return &ExampleServiceImpl{
		option: Option{},
	}
}

func (h *ExampleServiceImpl) RpcExample(ctx context.Context, request *api.ExampleRequest) (*api.ResponseVoid, error) {
	return &api.ResponseVoid{}, nil
}
