package service

import (
	"context"
	"github.com/MyGoFor/E-commerce/app/order/biz/dal/mysql"
	"github.com/MyGoFor/E-commerce/app/order/module"
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/cart"
	order "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/order"
	"log"
	"strconv"
)

type ListOrderService struct {
	ctx context.Context
} // NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	// Finish your business logic.
	var orders []*module.Order
	err = mysql.DB.Where("user_id = ?", req.UserId).Find(&orders).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var RespOrders []*order.Order
	for _, v := range orders {
		var Items []*order.OrderItem
		for _, k := range v.OrderItems {
			i := &order.OrderItem{
				Item: &cart.CartItem{
					ProductId: k.ProductId,
					Quantity:  k.Quantity,
				},
				Cost: k.Cost,
			}
			Items = append(Items, i)
		}
		atom, _ := strconv.Atoi(v.CreatedAt.String())
		o := &order.Order{
			OrderItems:   Items,
			OrderId:      v.OrderId,
			UserId:       v.UserId,
			UserCurrency: v.UserCurrency,
			Address: &order.Address{
				StreetAddress: v.Consignee.StreetAddress,
				City:          v.Consignee.City,
				State:         v.Consignee.State,
				Country:       v.Consignee.Country,
				ZipCode:       v.Consignee.ZipCode,
			},
			Email:     v.Consignee.Email,
			CreatedAt: int32(atom),
		}
		RespOrders = append(RespOrders, o)
	}
	return &order.ListOrderResp{Orders: RespOrders}, nil
}
