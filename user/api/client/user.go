package client

import (
	"context"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	user "github.com/jason4wy/mmicro/user/srv/proto/user"
)

type userKey struct {}

// FromContext retrieves the client from the Context
func UserFromContext(ctx context.Context) (user.UserService, bool) {
	c, ok := ctx.Value(userKey{}).(user.UserService)
	return c, ok
}

// Client returns a wrapper for the UserClient
func UserWrapper(service micro.Service) server.HandlerWrapper {
	client := user.NewUserService("io.github.jason4wy.srv.user", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, userKey{}, client)
			return fn(ctx, req, rsp)
		}
	}
}
