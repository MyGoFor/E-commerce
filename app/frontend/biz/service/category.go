package service

import (
	"context"
	category "github.com/MyGoFor/E-commerce/app/frontend/hertz_gen/frontend/category"
	"github.com/cloudwego/hertz/pkg/app"
)

type CategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCategoryService(Context context.Context, RequestContext *app.RequestContext) *CategoryService {
	return &CategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *CategoryService) Run(req *category.CategoryReq) (resp map[string]any, err error) {
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
			Name:        "02.1",
			Description: "02.1",
			Price:       1.999,
			Categories:  []string{"Two"},
			Picture:     "https://tuchuang.hch1212.online/img/021.webp",
		},
	}
	resp = map[string]any{
		"Title": "Category",
		"items": product,
	}
	return
}
