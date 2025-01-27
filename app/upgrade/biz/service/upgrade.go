package service

import (
	"context"
	"github.com/MyGoFor/E-commerce/app/casbin/middleware/ven"
	"github.com/MyGoFor/E-commerce/app/upgrade/biz/dal"
	upgrade "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/upgrade"
)

type UpgradeService struct {
	ctx context.Context
} // NewUpgradeService new UpgradeService
func NewUpgradeService(ctx context.Context) *UpgradeService {
	return &UpgradeService{ctx: ctx}
}

// Run create note info
func (s *UpgradeService) Run(req *upgrade.UpgradeReq) (resp *upgrade.Empty, err error) {
	// Finish your business logic.
	//资格检验略
	//升级
	err = ven.AddVendor(dal.E, req.Email)
	if err != nil {
		return nil, err
	}
	return &upgrade.Empty{}, nil
}
