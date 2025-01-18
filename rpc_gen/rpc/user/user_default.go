package user

import (
	"context"
	user "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Token(ctx context.Context, req *user.TokenReq, callOptions ...callopt.Option) (resp *user.Resp, err error) {
	resp, err = defaultClient.Token(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Token call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Refresh(ctx context.Context, req *user.RefreshTokenReq, callOptions ...callopt.Option) (resp *user.Resp, err error) {
	resp, err = defaultClient.Refresh(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Refresh call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
