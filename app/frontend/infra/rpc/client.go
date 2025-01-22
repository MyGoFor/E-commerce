package rpc

import (
	"github.com/MyGoFor/E-commerce/app/frontend/conf"
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"sync"
)

var (
	ProductClient productcatalogservice.Client
	UserClient    userservice.Client
	once          sync.Once
	err           error
	//registryAddr   string
	//commonSuite    client.Option
	//CurrentServiceName string
)

func Init() {
	once.Do(func() {
		initUserClient()
		initProductClient()
	})
}

func initUserClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	if err != nil {
		hlog.Fatal(err)
	}
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		hlog.Fatal(err)
	}
}

func initProductClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	if err != nil {
		hlog.Fatal(err)
	}
	ProductClient, err = productcatalogservice.NewClient("product", client.WithResolver(r))
	if err != nil {
		hlog.Fatal(err)
	}
}
