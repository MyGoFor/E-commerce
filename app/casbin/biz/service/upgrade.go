package service

import (
	"context"
	casbin "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/casbin"
)

type UpgradeService struct {
	ctx context.Context
} // NewUpgradeService new UpgradeService
func NewUpgradeService(ctx context.Context) *UpgradeService {
	return &UpgradeService{ctx: ctx}
}

// Run create note info
func (s *UpgradeService) Run(req *casbin.UpgradeReq) (resp *casbin.Empty, err error) {
	// Finish your business logic.

	return
}
