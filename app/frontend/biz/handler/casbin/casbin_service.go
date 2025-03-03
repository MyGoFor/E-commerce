package casbin

import (
	"context"
	"github.com/MyGoFor/E-commerce/app/frontend/biz/service"
	"github.com/MyGoFor/E-commerce/app/frontend/biz/utils"
	casbin "github.com/MyGoFor/E-commerce/app/frontend/hertz_gen/frontend/casbin"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// UpgradeOrOk .
// @router /casbin [POST]
func UpgradeOrOk(ctx context.Context, c *app.RequestContext) {
	var err error
	var req casbin.UpgradeOrOkReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	_, err = service.NewUpgradeOrOkService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "aboutOK", map[string]interface{}{"Title": "AboutOK"})
}
