package eino

import (
	"context"
	eino "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/eino"

	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/eino/einoservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() einoservice.Client
	Service() string
	SearchOrder(ctx context.Context, Req *eino.SearchOrderReq, callOptions ...callopt.Option) (r *eino.SearchOrderResp, err error)
	PlaceOrder(ctx context.Context, Req *eino.PlaceOrderReq, callOptions ...callopt.Option) (r *eino.PlaceOrderResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := einoservice.NewClient(dstService, opts...)
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
	kitexClient einoservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() einoservice.Client {
	return c.kitexClient
}

func (c *clientImpl) SearchOrder(ctx context.Context, Req *eino.SearchOrderReq, callOptions ...callopt.Option) (r *eino.SearchOrderResp, err error) {
	return c.kitexClient.SearchOrder(ctx, Req, callOptions...)
}

func (c *clientImpl) PlaceOrder(ctx context.Context, Req *eino.PlaceOrderReq, callOptions ...callopt.Option) (r *eino.PlaceOrderResp, err error) {
	return c.kitexClient.PlaceOrder(ctx, Req, callOptions...)
}
