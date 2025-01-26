package dal

import (
	"github.com/MyGoFor/E-commerce/app/casbin/db"
	"github.com/MyGoFor/E-commerce/app/upgrade/biz/dal/mysql"
	"github.com/MyGoFor/E-commerce/app/upgrade/biz/dal/redis"
	"github.com/casbin/casbin/v2"
)

var E *casbin.CachedEnforcer

func Init() {
	E = db.CasbinInit()
	redis.Init()
	mysql.Init()
}
