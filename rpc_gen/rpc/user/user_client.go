package user

import (
	"context"
	user "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/user"

	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() userservice.Client
	Service() string
	Token(ctx context.Context, Req *user.TokenReq, callOptions ...callopt.Option) (r *user.Resp, err error)
	Refresh(ctx context.Context, Req *user.RefreshTokenReq, callOptions ...callopt.Option) (r *user.Resp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := userservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient userservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() userservice.Client {
	return c.kitexClient
}

func (c *clientImpl) Token(ctx context.Context, Req *user.TokenReq, callOptions ...callopt.Option) (r *user.Resp, err error) {
	return c.kitexClient.Token(ctx, Req, callOptions...)
}

func (c *clientImpl) Refresh(ctx context.Context, Req *user.RefreshTokenReq, callOptions ...callopt.Option) (r *user.Resp, err error) {
	return c.kitexClient.Refresh(ctx, Req, callOptions...)
}
