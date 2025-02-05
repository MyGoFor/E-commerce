package db

import (
	"github.com/casbin/casbin/v2"
	"log"
)

func CasbinInit() *casbin.Enforcer {
	E, err := casbin.NewEnforcer("../../cas/model.pml", "../../cas/policy.csv")
	if err != nil {
		log.Println(err)
	}
	return E
}
