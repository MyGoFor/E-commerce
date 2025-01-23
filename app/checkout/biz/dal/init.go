package dal

import (
	"github.com/MyGoFor/E-commerce/app/checkout/biz/dal/mysql"
	"github.com/MyGoFor/E-commerce/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
