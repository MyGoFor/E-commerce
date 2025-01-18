package service

import (
	"context"
	"strconv"

	cart "github.com/MyGoFor/E-commerce/app/frontend/hertz_gen/frontend/cart"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *cart.Empty) (resp map[string]any, err error) {
	var items []map[string]string

	items = append(items, map[string]string{"Name": "02.0", "Description": "02.0", "Picture": "https://tuchuang.hch1212.online/img/023.webp", "Price": "1.99", "Qty": "1"})

	resp = map[string]any{
		"Title": "Cart",
		"items": items,
		"total": strconv.FormatFloat(float64(19), 'f', 2, 64),
	}
	return
}
