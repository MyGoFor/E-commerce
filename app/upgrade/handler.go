package main

import (
	"context"
	"github.com/MyGoFor/E-commerce/app/upgrade/biz/service"
	upgrade "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/upgrade"
)

// UpgradeServiceImpl implements the last service interface defined in the IDL.
type UpgradeServiceImpl struct{}

// Upgrade implements the UpgradeServiceImpl interface.
func (s *UpgradeServiceImpl) Upgrade(ctx context.Context, req *upgrade.UpgradeReq) (resp *upgrade.Empty, err error) {
	resp, err = service.NewUpgradeService(ctx).Run(req)

	return resp, err
}

// Ok implements the UpgradeServiceImpl interface.
func (s *UpgradeServiceImpl) Ok(ctx context.Context, req *upgrade.OkReq) (resp *upgrade.Empty, err error) {
	resp, err = service.NewOkService(ctx).Run(req)

	return resp, err
}
