package rpc

import (
	"context"
	"fmt"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/casbin/casbin/v2"
	"log"
)

func Check(e *casbin.Enforcer, ctx context.Context, obj, permission string) error {
	serviceName, _ := metainfo.GetValue(ctx, "SERVICE_NAME")

	log.Println(serviceName)

	// 使用 Casbin 进行权限检查
	allowed, err := e.Enforce(serviceName, obj, permission)
	if err != nil {
		return fmt.Errorf("error checking permission: %v", err)
	}
	if allowed {
		fmt.Printf("%s CAN %s %s\n", serviceName, obj, permission)
	} else {
		fmt.Printf("%s CANNOT %s %s\n", serviceName, obj, permission)
	}
	return nil
}
