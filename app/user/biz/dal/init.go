package dal

import (
	"github.com/MyGoFor/E-commerce/app/user/biz/dal/mysql"
	"github.com/MyGoFor/E-commerce/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
