package service

import (
	"context"

	product "github.com/MyGoFor/E-commerce/app/frontend/hertz_gen/frontend/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductService(Context context.Context, RequestContext *app.RequestContext) *GetProductService {
	return &GetProductService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductService) Run(req *product.ProductReq) (resp map[string]any, err error) {
	type Product struct {
		Id          uint32
		Name        string
		Description string
		Picture     string
		Price       float32
		Categories  []string
	}
	product := &Product{
		Id:          1,
		Name:        "02.0",
		Description: "02.0",
		Price:       1.99,
		Categories:  []string{"One"},
		Picture:     "https://tuchuang.hch1212.online/img/02.webp",
	}
	resp = map[string]any{
		"item": product,
	}
	return
}
