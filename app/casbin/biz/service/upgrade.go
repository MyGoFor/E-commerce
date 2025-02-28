package service

import (
	"context"
	"errors"
	"github.com/MyGoFor/E-commerce/app/casbin/conf/rbac"
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
	if req.Certificate == "" {
		err = errors.New("certificate is empty")
		return
	}
	// 新增一个用户role
	_, err = rbac.E.AddRoleForUser(req.Email, "business")
	if err != nil {
		return
	}
	err = rbac.E.SavePolicy()
	return
}
