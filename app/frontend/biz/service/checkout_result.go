package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"

	checkout "github.com/MyGoFor/E-commerce/app/frontend/hertz_gen/frontend/checkout"
	"github.com/cloudwego/hertz/pkg/app"
)

type CheckoutResultService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutResultService(Context context.Context, RequestContext *app.RequestContext) *CheckoutResultService {
	return &CheckoutResultService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutResultService) Run(req *checkout.Empty) (resp map[string]any, err error) {
	return utils.H{}, nil
}
