package eino

import (
	"context"
	eino "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/eino"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func SearchOrder(ctx context.Context, req *eino.SearchOrderReq, callOptions ...callopt.Option) (resp *eino.SearchOrderResp, err error) {
	resp, err = defaultClient.SearchOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "SearchOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func PlaceOrder(ctx context.Context, req *eino.PlaceOrderReq, callOptions ...callopt.Option) (resp *eino.PlaceOrderResp, err error) {
	resp, err = defaultClient.PlaceOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "PlaceOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
