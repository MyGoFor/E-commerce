package service

import (
	"context"
	"errors"
	"github.com/MyGoFor/E-commerce/app/casbin/middleware/ven"
	"github.com/MyGoFor/E-commerce/app/upgrade/biz/dal"
	upgrade "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/upgrade"
)

type OkService struct {
	ctx context.Context
} // NewOkService new OkService
func NewOkService(ctx context.Context) *OkService {
	return &OkService{ctx: ctx}
}

// Run create note info
func (s *OkService) Run(req *upgrade.OkReq) (resp *upgrade.Empty, err error) {
	// Finish your business logic.
	b := ven.Check(dal.E, req.Email)
	if !b {
		return nil, errors.New("not vendor")
	}
	return &upgrade.Empty{}, nil
}
