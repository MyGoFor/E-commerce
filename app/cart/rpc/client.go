package rpc

import (
	"github.com/MyGoFor/E-commerce/app/cart/utils"
	"github.com/MyGoFor/E-commerce/app/product/conf"
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"sync"
)

var (
	once          sync.Once
	ProductClient productcatalogservice.Client
)

func Init() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	utils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	utils.MustHandleError(err)
}
