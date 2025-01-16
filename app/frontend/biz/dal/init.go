package dal

import (
	"github.com/MyGoFor/E-commerce/app/frontend/biz/dal/mysql"
	"github.com/MyGoFor/E-commerce/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
