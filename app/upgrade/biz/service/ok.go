package service

import (
	"context"
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

	return
}
