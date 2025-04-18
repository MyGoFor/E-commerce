package rbac

import (
	"github.com/casbin/casbin/v2"
	"log"
)

var (
	E   *casbin.Enforcer
	err error
)

func InitCasbin() {
	E, err = casbin.NewEnforcer("conf/rbac/model.pml", "conf/rbac/policy.csv")
	if err != nil {
		log.Fatalf("NewEnforecer failed:%v\n", err)
	}
}
