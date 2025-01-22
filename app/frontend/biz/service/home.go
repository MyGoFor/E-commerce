package service

import (
	"context"
	home "github.com/MyGoFor/E-commerce/app/frontend/hertz_gen/frontend/home"
	"github.com/MyGoFor/E-commerce/app/frontend/infra/rpc"
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/klog"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *home.Empty) (resp map[string]any, err error) {
	ctx := h.Context
	p, err := rpc.ProductClient.ListProducts(ctx, &product.ListProductsReq{})
	if err != nil {
		klog.Error(err)
	}
	resp = map[string]any{
		"Items": p.Products,
		"Title": "E-commerce",
	}
	return
}
