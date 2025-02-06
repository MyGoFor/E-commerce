package main

import (
	"context"
	"github.com/MyGoFor/E-commerce/app/casbin/middleware/rpc"
	"github.com/MyGoFor/E-commerce/app/order/biz/dal"
	"github.com/MyGoFor/E-commerce/app/order/biz/service"
	order "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/order"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"log"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// PlaceOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) PlaceOrder(ctx context.Context, req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	//检查调用方权限
	serviceName, _ := metainfo.GetValue(ctx, "SERVICE_NAME")

	log.Println(serviceName)
	err = rpc.Check(dal.E, ctx, "place_order", "use")
	if err != nil {
		return nil, err
	}
	resp, err = service.NewPlaceOrderService(ctx).Run(req)

	return resp, err
}

// ListOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) ListOrder(ctx context.Context, req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	resp, err = service.NewListOrderService(ctx).Run(req)

	return resp, err
}

// MarkOrderPaid implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) MarkOrderPaid(ctx context.Context, req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	//检查调用方权限
	serviceName, _ := metainfo.GetValue(ctx, "service_name")

	log.Println(serviceName)
	err = rpc.Check(dal.E, ctx, "mark_order_paid", "use")
	if err != nil {
		return nil, err
	}
	resp, err = service.NewMarkOrderPaidService(ctx).Run(req)

	return resp, err
}
