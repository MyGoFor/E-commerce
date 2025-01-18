package service

import (
	"context"
	"errors"
	"github.com/MyGoFor/E-commerce/app/user/biz/dal/mysql"
	"github.com/MyGoFor/E-commerce/app/user/biz/model"
	user "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// Finish your business logic.
	if req.Email == "" || req.Password == "" || req.ConfirmPassword == "" {
		return nil, errors.New("email or password is empty")
	}
	if req.Password != req.ConfirmPassword {
		return nil, errors.New("password error")
	}

	passwordHashd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newUser := &model.User{
		Email:    req.Email,
		Password: string(passwordHashd),
	}

	err = model.Create(mysql.DB, s.ctx, newUser)
	if err != nil {
		return nil, err
	}

	return &user.RegisterResp{UserId: int32(newUser.ID)}, nil
}
