package service

import (
	"context"
	"strconv"

	checkout "github.com/MyGoFor/E-commerce/app/frontend/hertz_gen/frontend/checkout"
	"github.com/cloudwego/hertz/pkg/app"
)

type CheckoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutService(Context context.Context, RequestContext *app.RequestContext) *CheckoutService {
	return &CheckoutService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutService) Run(req *checkout.CheckoutReq) (resp map[string]any, err error) {
	var items []map[string]string

	items = append(items, map[string]string{
		"Name":    "02.0",
		"Price":   strconv.FormatFloat(float64(1.99), 'f', 2, 64),
		"Picture": "https://tuchuang.hch1212.online/img/021.webp",
		"Qty":     strconv.Itoa(int(1)),
	})

	resp = map[string]any{
		"Title":    "Checkout",
		"items":    items,
		"cart_num": len(items),
		"total":    strconv.FormatFloat(float64(19), 'f', 2, 64),
	}
	return
}
