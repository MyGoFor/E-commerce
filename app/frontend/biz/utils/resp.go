package utils

import (
	"context"
	"github.com/hertz-contrib/sessions"

	"github.com/cloudwego/hertz/pkg/app"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.String(code, err.Error())
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, data)
}

// 获取一些上下文信息
func WarpResponse(ctx context.Context, c *app.RequestContext, content map[string]any) map[string]any {
	session := sessions.Default(c)
	userId := session.Get("user_id")
	content["user_id"] = userId
	return content
}
