package service

import (
	"context"
	"github.com/hertz-contrib/sessions"

	auth "github.com/MyGoFor/E-commerce/app/frontend/hertz_gen/frontend/auth"
	"github.com/cloudwego/hertz/pkg/app"
)

type LogoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLogoutService(Context context.Context, RequestContext *app.RequestContext) *LogoutService {
	return &LogoutService{RequestContext: RequestContext, Context: Context}
}

func (h *LogoutService) Run(req *auth.Empty) (resp *auth.Empty, err error) {
	session := sessions.Default(h.RequestContext)
	session.Clear()
	err = session.Save()
	if err != nil {
		return nil, err
	}
	return
}
