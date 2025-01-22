package service

import (
	"context"
	"errors"
	"github.com/MyGoFor/E-commerce/app/order/biz/dal/mysql"
	"github.com/MyGoFor/E-commerce/app/order/module"
	order "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/order"
)

type MarkOrderPaidService struct {
	ctx context.Context
} // NewMarkOrderPaidService new MarkOrderPaidService
func NewMarkOrderPaidService(ctx context.Context) *MarkOrderPaidService {
	return &MarkOrderPaidService{ctx: ctx}
}

// Run create note info
func (s *MarkOrderPaidService) Run(req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	// Finish your business logic.
	if req.OrderId == "" && req.UserId == 0 {
		return nil, errors.New("nil")
	}
	var o module.Order
	err = mysql.DB.Where("order_id = ? AND user_id = ?", req.OrderId, req.UserId).First(&o).Error
	if err != nil {
		return nil, err
	}
	err = mysql.DB.Model(&o).Update("order_state", module.OrderStatePaid).Error
	if err != nil {
		return nil, err
	}
	resp = &order.MarkOrderPaidResp{}
	return resp, nil
}
