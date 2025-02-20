package main

import (
	"context"
	"github.com/MyGoFor/E-commerce/app/casbin/biz/service"
	casbin "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/casbin"
)

// CasbinServiceImpl implements the last service interface defined in the IDL.
type CasbinServiceImpl struct{}

// Upgrade implements the CasbinServiceImpl interface.
func (s *CasbinServiceImpl) Upgrade(ctx context.Context, req *casbin.UpgradeReq) (resp *casbin.Empty, err error) {
	resp, err = service.NewUpgradeService(ctx).Run(req)

	return resp, err
}

// Ok implements the CasbinServiceImpl interface.
func (s *CasbinServiceImpl) Ok(ctx context.Context, req *casbin.OkReq) (resp *casbin.Empty, err error) {
	resp, err = service.NewOkService(ctx).Run(req)

	return resp, err
}
