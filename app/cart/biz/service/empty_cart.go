package service

import (
	"context"
	"github.com/MyGoFor/E-commerce/app/cart/biz/dal"
	"github.com/MyGoFor/E-commerce/app/cart/biz/dal/mysql"
	"github.com/MyGoFor/E-commerce/app/cart/biz/model"
	"github.com/MyGoFor/E-commerce/app/casbin/middleware/rpc"
	cart "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type EmptyCartService struct {
	ctx context.Context
} // NewEmptyCartService new EmptyCartService
func NewEmptyCartService(ctx context.Context) *EmptyCartService {
	return &EmptyCartService{ctx: ctx}
}

// Run create note info
func (s *EmptyCartService) Run(req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	//检查调用方权限
	err = rpc.Check(dal.E, s.ctx, "empty_cart", "write")
	if err != nil {
		return nil, err
	}
	// Finish your business logic.
	err = model.EmptyCart(mysql.DB, s.ctx, req.GetUserId())
	if err != nil {
		return &cart.EmptyCartResp{}, kerrors.NewBizStatusError(50001, "empty cart error")
	}

	return &cart.EmptyCartResp{}, nil
}
