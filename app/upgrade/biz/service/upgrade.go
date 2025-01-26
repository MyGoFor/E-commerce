package service

import (
	"context"
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
	return
}
