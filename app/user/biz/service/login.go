package service

import (
	"context"
	"github.com/MyGoFor/E-commerce/app/user/biz/dal/mysql"
	"github.com/MyGoFor/E-commerce/app/user/biz/dal/redis"
	"github.com/MyGoFor/E-commerce/app/user/biz/model"
	user "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.
	klog.Infof("LoginReq:%+v", req)
	userRow, err := model.NewCacheUserQuery(model.NewUserQuery(s.ctx, mysql.DB), redis.RedisClient).GetByEmail(req.Email)
	if err != nil {
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(userRow.Password), []byte(req.Password))
	if err != nil {
		return
	}
	return &user.LoginResp{UserId: int32(userRow.ID)}, nil
}
