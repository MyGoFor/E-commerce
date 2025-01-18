// Code generated by Kitex v0.9.1. DO NOT EDIT.

package userservice

import (
	"context"
	user "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/user"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Token(ctx context.Context, Req *user.TokenReq, callOptions ...callopt.Option) (r *user.Resp, err error)
	Refresh(ctx context.Context, Req *user.RefreshTokenReq, callOptions ...callopt.Option) (r *user.Resp, err error)
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
	return &kUserServiceClient{
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

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) Token(ctx context.Context, Req *user.TokenReq, callOptions ...callopt.Option) (r *user.Resp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Token(ctx, Req)
}

func (p *kUserServiceClient) Refresh(ctx context.Context, Req *user.RefreshTokenReq, callOptions ...callopt.Option) (r *user.Resp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Refresh(ctx, Req)
}
