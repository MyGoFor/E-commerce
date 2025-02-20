package dal

import (
	"github.com/MyGoFor/E-commerce/app/casbin/biz/dal/mysql"
	"github.com/MyGoFor/E-commerce/app/casbin/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
