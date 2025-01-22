package service

import (
	"context"
	"time"

	order "github.com/MyGoFor/E-commerce/app/frontend/hertz_gen/frontend/order"
	"github.com/cloudwego/hertz/pkg/app"
)

type OrderListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderListService(Context context.Context, RequestContext *app.RequestContext) *OrderListService {
	return &OrderListService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderListService) Run(req *order.Empty) (resp map[string]any, err error) {
	type Consignee struct {
		Email         string
		StreetAddress string
		City          string
		State         string
		Country       string
		ZipCode       int32
	}

	type OrderItem struct {
		ProductId   uint32
		ProductName string
		Picture     string
		Qty         uint32
		Cost        float32
	}

	type Order struct {
		Consignee   Consignee
		OrderId     string
		CreatedDate string
		OrderState  string
		Cost        float32
		Items       []OrderItem
	}
	var orders []*Order

	var items []OrderItem
	items = append(items, OrderItem{
		ProductId:   1,
		Qty:         uint32(3),
		ProductName: "02.0",
		Picture:     "https://tuchuang.hch1212.online/img/023.webp",
		Cost:        1.99,
	})

	orders = append(orders, &Order{
		Cost:        10,
		Items:       items,
		CreatedDate: time.Now().Format("2006-01-02 15:04:05"),
		OrderId:     "1",
		Consignee:   Consignee{Email: "12@qq.com"},
	})

	resp = map[string]any{
		"Title": "Order",
		"order": orders,
	}
	return
}
