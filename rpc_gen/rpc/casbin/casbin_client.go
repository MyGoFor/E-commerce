package casbin

import (
	"context"
	casbin "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/casbin"

	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/casbin/casbinservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() casbinservice.Client
	Service() string
	Upgrade(ctx context.Context, Req *casbin.UpgradeReq, callOptions ...callopt.Option) (r *casbin.Empty, err error)
	Ok(ctx context.Context, Req *casbin.OkReq, callOptions ...callopt.Option) (r *casbin.Empty, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := casbinservice.NewClient(dstService, opts...)
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
	kitexClient casbinservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() casbinservice.Client {
	return c.kitexClient
}

func (c *clientImpl) Upgrade(ctx context.Context, Req *casbin.UpgradeReq, callOptions ...callopt.Option) (r *casbin.Empty, err error) {
	return c.kitexClient.Upgrade(ctx, Req, callOptions...)
}

func (c *clientImpl) Ok(ctx context.Context, Req *casbin.OkReq, callOptions ...callopt.Option) (r *casbin.Empty, err error) {
	return c.kitexClient.Ok(ctx, Req, callOptions...)
}
