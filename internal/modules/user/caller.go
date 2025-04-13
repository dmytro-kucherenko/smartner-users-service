package user

import (
	"context"

	usera "github.com/dmytro-kucherenko/smartner-contracts-package/gen/go/user"
	"github.com/dmytro-kucherenko/smartner-users-service/internal/modules/user/services"
	adapter "github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/adapters/grpc"
	"google.golang.org/grpc"
)

type Caller struct {
	usera.UnimplementedUserServer
	service *services.Main
}

func NewCaller(service *services.Main) *Caller {
	return &Caller{service: service}
}

func (caller *Caller) Init(server *grpc.Server) {
	usera.RegisterUserServer(server, caller)
}

func (caller *Caller) Config() adapter.CallerConfig {
	return adapter.CallerConfig{
		usera.User_Get_FullMethodName: {Public: false},
	}
}

func (caller *Caller) Get(ctx context.Context, req *usera.GetRequest) (res *usera.GetResponse, err error) {
	return adapter.HandleProcedure(caller.service.Get, ctx, req, res)
}
