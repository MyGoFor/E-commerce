package service

import (
	"context"
	casbin "github.com/MyGoFor/E-commerce/app/frontend/hertz_gen/frontend/casbin"
	"github.com/MyGoFor/E-commerce/app/frontend/infra/rpc"
	rpccasbin "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/casbin"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpgradeOrOkService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpgradeOrOkService(Context context.Context, RequestContext *app.RequestContext) *UpgradeOrOkService {
	return &UpgradeOrOkService{RequestContext: RequestContext, Context: Context}
}

func (h *UpgradeOrOkService) Run(req *casbin.UpgradeOrOkReq) (resp *casbin.Empty, err error) {
	if req.Certificate == "" {
		// 验证逻辑
		_, err = rpc.CasbinClient.Ok(h.Context, &rpccasbin.OkReq{Sub: req.Email, Obj: "/about", Act: "GET"})
	} else {
		// 升级逻辑
		_, err = rpc.CasbinClient.Upgrade(h.Context, &rpccasbin.UpgradeReq{Email: req.Email, Certificate: req.Certificate})
	}
	return
}
