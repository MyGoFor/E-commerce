package service

import (
	"context"
	user "github.com/MyGoFor/E-commerce/app/user/kitex_gen/user"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.Resp, err error) {
	// Finish your business logic.

	return
}
