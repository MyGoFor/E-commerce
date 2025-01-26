package rpc

import (
	"context"
	"fmt"
	"github.com/casbin/casbin/v2"
)

func Check(e *casbin.CachedEnforcer, ctx context.Context, obj, permission string) error {
	serviceName, _ := ctx.Value("service_name").(string)

	// 使用 Casbin 进行权限检查
	allowed, err := e.Enforce(serviceName, obj, permission)
	if err != nil {
		return fmt.Errorf("error checking permission: %v", err)
	}
	if !allowed {
		return fmt.Errorf("permission denied for %s to call %s", serviceName, obj)
	}
	return nil
}
