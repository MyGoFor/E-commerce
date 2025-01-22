package service

import (
	"context"
	"github.com/MyGoFor/E-commerce/app/order/biz/dal/mysql"
	"github.com/MyGoFor/E-commerce/app/order/module"
	order "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/order"
	"github.com/google/uuid"
	"log"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	// Finish your business logic.
	Uuid, _ := uuid.NewUUID()
	var items []module.OrderItem
	for _, v := range req.OrderItems {
		item := module.OrderItem{
			ProductId:    v.Item.ProductId,
			OrderIdRefer: Uuid.String(),
			Quantity:     v.Item.Quantity,
			Cost:         v.Cost,
		}
		items = append(items, item)
	}
	userOrder := module.Order{
		OrderId:      Uuid.String(),
		UserId:       req.UserId,
		UserCurrency: req.UserCurrency,
		Consignee: module.Consignee{
			Email:         req.Email,
			StreetAddress: req.Address.StreetAddress,
			City:          req.Address.City,
			State:         req.Address.State,
			Country:       req.Address.Country,
			ZipCode:       req.Address.ZipCode,
		},
		OrderItems: items,
		OrderState: module.OrderStatePlaced,
	}
	err = mysql.DB.Create(&userOrder).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &order.PlaceOrderResp{Order: &order.OrderResult{OrderId: userOrder.OrderId}}, nil
}
