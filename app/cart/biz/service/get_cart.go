package service

import (
	"context"
	"github.com/MyGoFor/E-commerce/app/cart/biz/model"
	"github.com/MyGoFor/E-commerce/app/product/biz/dal/mysql"
	cart "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	// Finish your business logic.
	list, err := model.GetCartByUserID(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50002, err.Error())
	}
	var items []*cart.CartItem
	for _, item := range list {
		items = append(items, &cart.CartItem{
			ProductId: item.ProductID,
			Quantity:  int32(item.Qty),
		})
	}
	return &cart.GetCartResp{Cart: &cart.Cart{
		UserId: req.GetUserId(),
		Items:  items,
	}}, nil
}
