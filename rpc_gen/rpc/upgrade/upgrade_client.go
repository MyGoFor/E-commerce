package upgrade

import (
	"context"
	upgrade "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/upgrade"

	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/upgrade/upgradeservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() upgradeservice.Client
	Service() string
	Upgrade(ctx context.Context, Req *upgrade.UpgradeReq, callOptions ...callopt.Option) (r *upgrade.Empty, err error)
	Ok(ctx context.Context, Req *upgrade.OkReq, callOptions ...callopt.Option) (r *upgrade.Empty, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := upgradeservice.NewClient(dstService, opts...)
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
	kitexClient upgradeservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() upgradeservice.Client {
	return c.kitexClient
}

func (c *clientImpl) Upgrade(ctx context.Context, Req *upgrade.UpgradeReq, callOptions ...callopt.Option) (r *upgrade.Empty, err error) {
	return c.kitexClient.Upgrade(ctx, Req, callOptions...)
}

func (c *clientImpl) Ok(ctx context.Context, Req *upgrade.OkReq, callOptions ...callopt.Option) (r *upgrade.Empty, err error) {
	return c.kitexClient.Ok(ctx, Req, callOptions...)
}
