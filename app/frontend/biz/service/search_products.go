package service

import (
	"context"

	product "github.com/MyGoFor/E-commerce/app/frontend/hertz_gen/frontend/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type SearchProductsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProductsService(Context context.Context, RequestContext *app.RequestContext) *SearchProductsService {
	return &SearchProductsService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProductsService) Run(req *product.SearchProductsReq) (resp map[string]any, err error) {
	type Product struct {
		Id          uint32
		Name        string
		Description string
		Picture     string
		Price       float32
		Categories  []string
	}
	product := []*Product{{
		Id:          1,
		Name:        "02.0",
		Description: "02.0",
		Price:       1.99,
		Categories:  []string{"One"},
		Picture:     "https://tuchuang.hch1212.online/img/02.webp",
	},
		{
			Id:          2,
			Name:        "02.0",
			Description: "02.0",
			Price:       1.999,
			Categories:  []string{"Two"},
			Picture:     "https://tuchuang.hch1212.online/img/02.webp",
		},
	}
	resp = map[string]any{
		"items": product,
		"q":     req.Q,
	}
	return
}
