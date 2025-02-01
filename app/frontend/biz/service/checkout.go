package service

import (
	"context"
	"github.com/MyGoFor/E-commerce/app/frontend/infra/rpc"
	rpccart "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/cart"
	rpcproduct "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/common/utils"
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
	userId := uint32(h.Context.Value("user_id").(int32))

	carts, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartReq{UserId: userId})
	if err != nil {
		return nil, err
	}
	var total float32
	for _, v := range carts.Cart.Items {
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: v.ProductId})
		if err != nil {
			return nil, err
		}
		if productResp.Product == nil {
			continue
		}
		p := productResp.Product
		items = append(items, map[string]string{
			"Name":    p.Name,
			"Price":   strconv.FormatFloat(float64(p.Price), 'f', 2, 64),
			"Picture": p.Picture,
			"Qty":     strconv.Itoa(int(v.Quantity)),
		})
		total += float32(v.Quantity) * p.Price
	}

	return utils.H{
		"Title":    "Checkout",
		"items":    items,
		"cart_num": len(items),
		"total":    strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}
