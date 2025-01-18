package order

import (
	"context"

	"github.com/MyGoFor/E-commerce/app/frontend/biz/service"
	"github.com/MyGoFor/E-commerce/app/frontend/biz/utils"
	order "github.com/MyGoFor/E-commerce/app/frontend/hertz_gen/frontend/order"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// OrderList .
// @router /order [GET]
func OrderList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewOrderListService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "order", utils.WarpResponse(ctx, c, resp))
}
