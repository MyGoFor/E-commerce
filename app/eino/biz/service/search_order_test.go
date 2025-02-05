package service

import (
	"context"
	eino "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/eino"
	"testing"
)

func TestSearchOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSearchOrderService(ctx)
	// init req and assert value

	req := &eino.SearchOrderReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
