package service

import (
	"context"
	"errors"
	"github.com/MyGoFor/E-commerce/app/casbin/conf/rbac"
	casbin "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/casbin"
)

type OkService struct {
	ctx context.Context
} // NewOkService new OkService
func NewOkService(ctx context.Context) *OkService {
	return &OkService{ctx: ctx}
}

// Run create note info
func (s *OkService) Run(req *casbin.OkReq) (resp *casbin.Empty, err error) {
	// Finish your business logic.
	ok, _ := rbac.E.Enforce(req.Sub, req.Obj, req.Act)
	if !ok {
		err = errors.New("不是商家")
	}
	return
}
