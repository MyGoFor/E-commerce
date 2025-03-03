package service

import (
	"context"
	"github.com/MyGoFor/E-commerce/app/eino/biz/AiModel"
	eino "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/eino"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *eino.PlaceOrderReq) (resp string, err error) {
	s.ctx = context.WithValue(s.ctx, "SERVICE_NAME", "ai")
	// Finish your business logic.
	resp, err = AiModel.PlaceModel(s.ctx, req.Question, req.Uid)
	if err != nil {
		return "", err
	}
	return resp, nil
}
