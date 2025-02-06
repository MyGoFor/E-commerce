package main

import (
	"context"
	"github.com/MyGoFor/E-commerce/app/cart/biz/dal"
	"github.com/MyGoFor/E-commerce/app/cart/biz/service"
	"github.com/MyGoFor/E-commerce/app/casbin/middleware/rpc"
	cart "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/cart"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"log"
)

// CartServiceImpl implements the last service interface defined in the IDL.
type CartServiceImpl struct{}

// AddItem implements the CartServiceImpl interface.
func (s *CartServiceImpl) AddItem(ctx context.Context, req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	resp, err = service.NewAddItemService(ctx).Run(req)

	return resp, err
}

// GetCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) GetCart(ctx context.Context, req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	resp, err = service.NewGetCartService(ctx).Run(req)

	return resp, err
}

// EmptyCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) EmptyCart(ctx context.Context, req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	//检查调用方权限
	serviceName, _ := metainfo.GetValue(ctx, "SERVICE_NAME")

	log.Println(serviceName)
	err = rpc.Check(dal.E, ctx, "empty_cart", "use")
	if err != nil {
		return nil, err
	}
	resp, err = service.NewEmptyCartService(ctx).Run(req)

	return resp, err
}
