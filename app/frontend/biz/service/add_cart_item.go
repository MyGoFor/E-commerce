package service

import (
	"context"
	cart "github.com/MyGoFor/E-commerce/app/frontend/hertz_gen/frontend/cart"
	"github.com/MyGoFor/E-commerce/app/frontend/infra/rpc"
	rpccart "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/hertz/pkg/app"
)

type AddCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartItemService(Context context.Context, RequestContext *app.RequestContext) *AddCartItemService {
	return &AddCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartItemService) Run(req *cart.AddCartReq) (resp *cart.Empty, err error) {
	_, err = rpc.CartClient.AddItem(h.Context, &rpccart.AddItemReq{
		//UserId: frontendutils.GetUserIdFromCtx(h.Context),
		UserId: uint32(h.Context.Value("user_id").(int32)),
		Item: &rpccart.CartItem{
			ProductId: req.ProductId,
			Quantity:  req.ProductNum,
		},
	})
	return
}
