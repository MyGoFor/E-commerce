package main

import (
	"context"
	"github.com/MyGoFor/E-commerce/app/casbin/middleware/rpc"
	"github.com/MyGoFor/E-commerce/app/payment/biz/dal"
	"github.com/MyGoFor/E-commerce/app/payment/biz/service"
	payment "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/payment"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"log"
)

// PaymentServiceImpl implements the last service interface defined in the IDL.
type PaymentServiceImpl struct{}

// Charge implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) Charge(ctx context.Context, req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	//检查调用方权限
	serviceName, _ := metainfo.GetValue(ctx, "SERVICE_NAME")

	log.Println(serviceName)
	err = rpc.Check(dal.E, ctx, "charge", "use")
	if err != nil {
		return nil, err
	}
	resp, err = service.NewChargeService(ctx).Run(req)

	return resp, err
}
