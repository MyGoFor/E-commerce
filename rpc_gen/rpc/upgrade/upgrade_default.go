package upgrade

import (
	"context"
	upgrade "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/upgrade"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Upgrade(ctx context.Context, req *upgrade.UpgradeReq, callOptions ...callopt.Option) (resp *upgrade.Empty, err error) {
	resp, err = defaultClient.Upgrade(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Upgrade call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Ok(ctx context.Context, req *upgrade.OkReq, callOptions ...callopt.Option) (resp *upgrade.Empty, err error) {
	resp, err = defaultClient.Ok(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Ok call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
