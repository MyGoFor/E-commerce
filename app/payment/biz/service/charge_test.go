package service

import (
	"context"
	payment "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/payment"
	"testing"
)

func TestCharge_Run(t *testing.T) {
	ctx := context.Background()
	s := NewChargeService(ctx)
	// init req and assert value
	req := &payment.ChargeReq{
		Amount:     1000,
		OrderId:    "1",
		UserId:     1,
		CreditCard: nil,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
