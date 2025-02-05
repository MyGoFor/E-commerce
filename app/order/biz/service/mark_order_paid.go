package service

import (
	"context"
	"fmt"
	"github.com/MyGoFor/E-commerce/app/casbin/middleware/rpc"
	"github.com/MyGoFor/E-commerce/app/order/biz/dal/mysql"
	"github.com/MyGoFor/E-commerce/app/order/biz/model"
	"github.com/MyGoFor/E-commerce/app/payment/biz/dal"
	order "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/klog"
)

type MarkOrderPaidService struct {
	ctx context.Context
} // NewMarkOrderPaidService new MarkOrderPaidService
func NewMarkOrderPaidService(ctx context.Context) *MarkOrderPaidService {
	return &MarkOrderPaidService{ctx: ctx}
}

// Run create note info
func (s *MarkOrderPaidService) Run(req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	//检查调用方权限
	err = rpc.Check(dal.E, s.ctx, "mark_order_paid", "write")
	if err != nil {
		return nil, err
	}
	// Finish your business logic.
	if req.UserId == 0 || req.OrderId == "" {
		err = fmt.Errorf("user_id or order_id can not be empty")
		return
	}
	_, err = model.GetOrder(mysql.DB, s.ctx, req.UserId, req.OrderId)
	if err != nil {
		klog.Errorf("model.ListOrder.err:%v", err)
		return nil, err
	}
	err = model.UpdateOrderState(mysql.DB, s.ctx, req.UserId, req.OrderId, model.OrderStatePaid)
	if err != nil {
		klog.Errorf("model.ListOrder.err:%v", err)
		return nil, err
	}
	resp = &order.MarkOrderPaidResp{}
	return
}
