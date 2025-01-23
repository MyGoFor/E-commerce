package dal

import (
	"github.com/MyGoFor/E-commerce/app/order/biz/dal/mysql"
	"github.com/MyGoFor/E-commerce/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
