package service

import (
	"context"
	product "github.com/MyGoFor/E-commerce/app/frontend/hertz_gen/frontend/product"
	"github.com/MyGoFor/E-commerce/app/frontend/infra/rpc"
	rpcproduct "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type AddProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddProductService(Context context.Context, RequestContext *app.RequestContext) *AddProductService {
	return &AddProductService{RequestContext: RequestContext, Context: Context}
}

func (h *AddProductService) Run(req *product.AddReq) (resp *product.Empty, err error) {
	_, err = rpc.ProductClient.AddProduct(h.Context, &rpcproduct.AddProductReq{
		Product: &rpcproduct.Product{
			Name:        req.Name,
			Description: req.Description,
			Price:       req.Price,
			Picture:     req.Picture,
			Categories:  req.Categories,
		},
	})
	return
}
