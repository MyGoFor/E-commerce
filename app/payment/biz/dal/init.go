package dal

import (
	"github.com/MyGoFor/E-commerce/app/payment/biz/dal/mysql"
	"github.com/MyGoFor/E-commerce/app/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
