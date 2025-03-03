package ai

import (
	"context"
	"fmt"
	"github.com/MyGoFor/E-commerce/app/frontend/biz/service"
	"github.com/MyGoFor/E-commerce/app/frontend/biz/utils"
	ai "github.com/MyGoFor/E-commerce/app/frontend/hertz_gen/frontend/ai"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// ToCart .
// @router /eino [POST]
func ToCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req ai.AIReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	if req.Content == "" {
		utils.SendErrResponse(ctx, c, consts.StatusOK, fmt.Errorf("content is empty"))
		return
	}

	resp, err := service.NewToCartService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "aiOK", map[string]interface{}{
		"Title":   "AIOK",
		"Content": resp,
	})
}
