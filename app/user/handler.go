package main

import (
	"context"
	"github.com/MyGoFor/E-commerce/app/user/biz/service"
	"github.com/MyGoFor/E-commerce/app/user/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.Resp, err error) {
	resp, err = service.NewRegisterService(ctx).Run(req)

	return resp, err
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.Resp, err error) {
	resp, err = service.NewLoginService(ctx).Run(req)

	return resp, err
}
