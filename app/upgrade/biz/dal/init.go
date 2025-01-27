package dal

import (
	"github.com/MyGoFor/E-commerce/app/casbin/db"
	"github.com/casbin/casbin/v2"
)

var E *casbin.CachedEnforcer

func Init() {
	E = db.CasbinInit()
}
