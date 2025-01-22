package dal

import (
	"github.com/MyGoFor/E-commerce/app/product/biz/dal/mysql"
	"github.com/MyGoFor/E-commerce/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
