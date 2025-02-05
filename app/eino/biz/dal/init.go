package dal

import (
	"github.com/MyGoFor/E-commerce/app/eino/biz/dal/mysql"
	"github.com/MyGoFor/E-commerce/app/eino/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
