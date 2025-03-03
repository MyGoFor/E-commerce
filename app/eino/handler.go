package main

import (
	"context"
	"github.com/MyGoFor/E-commerce/app/eino/biz/service"
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/eino"
)

// EinoServiceImpl implements the last service interface defined in the IDL.
type EinoServiceImpl struct{}

// SearchOrder implements the EinoServiceImpl interface.
func (s *EinoServiceImpl) SearchOrder(ctx context.Context, req *eino.SearchOrderReq) (resp *eino.SearchOrderResp, err error) {
	resp, err = service.NewSearchOrderService(ctx).Run(req)

	return resp, err
}

// PlaceOrder implements the EinoServiceImpl interface.
func (s *EinoServiceImpl) PlaceOrder(ctx context.Context, req *eino.PlaceOrderReq) (resp *eino.Empty, err error) {
	resp, err = service.NewPlaceOrderService(ctx).Run(req)

	return resp, err
}
