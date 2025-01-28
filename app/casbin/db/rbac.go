package db

import (
	"fmt"
	"github.com/MyGoFor/E-commerce/app/casbin/conf"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadpter "github.com/casbin/gorm-adapter/v3"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&gormadpter.CasbinRule{})
	if err != nil {
		log.Println(err)
	}
}

func CasbinInit() *casbin.CachedEnforcer {
	_ = godotenv.Load("./.env")
	Init()
	a, _ := gormadpter.NewAdapterByDB(DB)
	m, err := model.NewModelFromFile("../../pml/model.pml")
	if err != nil {
		log.Fatalln(err)
	}
	E, _ := casbin.NewCachedEnforcer(m, a)
	E.SetExpireTime(60 * 60)
	_ = E.LoadPolicy()
	return E
}
