package service

import (
	"context"
	eino "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/eino"
)

type SearchOrderService struct {
	ctx context.Context
} // NewSearchOrderService new SearchOrderService
func NewSearchOrderService(ctx context.Context) *SearchOrderService {
	return &SearchOrderService{ctx: ctx}
}

// Run create note info
func (s *SearchOrderService) Run(req *eino.SearchOrderReq) (resp *eino.SearchOrderResp, err error) {
	// Finish your business logic.

	return
}
