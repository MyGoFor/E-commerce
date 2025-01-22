// Code generated by Kitex v0.9.1. DO NOT EDIT.

package productcatalogservice

import (
	"context"
	"github.com/MyGoFor/E-commerce/rpc_gen/product"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	AddProduct(ctx context.Context, Req *product.AddProductReq, callOptions ...callopt.Option) (r *product.AddProductResp, err error)
	ListProducts(ctx context.Context, Req *product.ListProductsReq, callOptions ...callopt.Option) (r *product.ListProductsResp, err error)
	GetProduct(ctx context.Context, Req *product.GetProductReq, callOptions ...callopt.Option) (r *product.GetProductResp, err error)
	SearchProducts(ctx context.Context, Req *product.SearchProductsReq, callOptions ...callopt.Option) (r *product.SearchProductsResp, err error)
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
	return &kProductCatalogServiceClient{
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

type kProductCatalogServiceClient struct {
	*kClient
}

func (p *kProductCatalogServiceClient) AddProduct(ctx context.Context, Req *product.AddProductReq, callOptions ...callopt.Option) (r *product.AddProductResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddProduct(ctx, Req)
}

func (p *kProductCatalogServiceClient) ListProducts(ctx context.Context, Req *product.ListProductsReq, callOptions ...callopt.Option) (r *product.ListProductsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListProducts(ctx, Req)
}

func (p *kProductCatalogServiceClient) GetProduct(ctx context.Context, Req *product.GetProductReq, callOptions ...callopt.Option) (r *product.GetProductResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetProduct(ctx, Req)
}

func (p *kProductCatalogServiceClient) SearchProducts(ctx context.Context, Req *product.SearchProductsReq, callOptions ...callopt.Option) (r *product.SearchProductsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SearchProducts(ctx, Req)
}
