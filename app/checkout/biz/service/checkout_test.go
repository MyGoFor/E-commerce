package service

import (
	"context"
	checkout "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/checkout"
	"testing"
)

func TestCheckout_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCheckoutService(ctx)
	// init req and assert value

	req := &checkout.CheckoutReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
