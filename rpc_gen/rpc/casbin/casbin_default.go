package casbin

import (
	"context"
	casbin "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/casbin"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Upgrade(ctx context.Context, req *casbin.UpgradeReq, callOptions ...callopt.Option) (resp *casbin.Empty, err error) {
	resp, err = defaultClient.Upgrade(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Upgrade call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Ok(ctx context.Context, req *casbin.OkReq, callOptions ...callopt.Option) (resp *casbin.Empty, err error) {
	resp, err = defaultClient.Ok(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Ok call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
