package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		userId := session.Get("user_id")
		if userId == nil {
			c.Redirect(302, []byte("/sign-in"))
			c.Abort()
			return
		}
		ctx = context.WithValue(ctx, "user_id", userId)
		c.Next(ctx)
	}
}
