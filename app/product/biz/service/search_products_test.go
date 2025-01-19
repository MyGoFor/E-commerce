package service

import (
	"context"
	product "github.com/MyGoFor/E-commerce/rpc_gen/product"
	"testing"
)

func TestSearchProducts_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSearchProductsService(ctx)
	// init req and assert value

	req := &product.SearchProductsReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
