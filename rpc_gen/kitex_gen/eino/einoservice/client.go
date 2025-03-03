// Code generated by Kitex v0.9.1. DO NOT EDIT.

package einoservice

import (
	"context"
	eino "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/eino"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	SearchOrder(ctx context.Context, Req *eino.SearchOrderReq, callOptions ...callopt.Option) (r *eino.SearchOrderResp, err error)
	PlaceOrder(ctx context.Context, Req *eino.PlaceOrderReq, callOptions ...callopt.Option) (r *eino.PlaceOrderResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kEinoServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kEinoServiceClient struct {
	*kClient
}

func (p *kEinoServiceClient) SearchOrder(ctx context.Context, Req *eino.SearchOrderReq, callOptions ...callopt.Option) (r *eino.SearchOrderResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SearchOrder(ctx, Req)
}

func (p *kEinoServiceClient) PlaceOrder(ctx context.Context, Req *eino.PlaceOrderReq, callOptions ...callopt.Option) (r *eino.PlaceOrderResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PlaceOrder(ctx, Req)
}
