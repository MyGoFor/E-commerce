package service

import (
	"context"

	checkout "github.com/MyGoFor/E-commerce/app/frontend/hertz_gen/frontend/checkout"
	"github.com/cloudwego/hertz/pkg/app"
)

type CheckoutWaitingService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutWaitingService(Context context.Context, RequestContext *app.RequestContext) *CheckoutWaitingService {
	return &CheckoutWaitingService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutWaitingService) Run(req *checkout.Empty) (resp map[string]any, err error) {
	resp = map[string]any{
		"Title":    "waiting",
		"redirect": "/checkout/result",
	}
	return
}
