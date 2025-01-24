package dal

import (
	"github.com/MyGoFor/E-commerce/app/email/biz/dal/mysql"
	"github.com/MyGoFor/E-commerce/app/email/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
