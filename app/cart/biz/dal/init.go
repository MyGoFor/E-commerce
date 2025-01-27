package dal

import (
	"github.com/MyGoFor/E-commerce/app/cart/biz/dal/mysql"
	"github.com/MyGoFor/E-commerce/app/cart/biz/dal/redis"
	"github.com/MyGoFor/E-commerce/app/casbin/db"
	"github.com/casbin/casbin/v2"
)

var E *casbin.CachedEnforcer

func Init() {
	E = db.CasbinInit()
	redis.Init()
	mysql.Init()
}
