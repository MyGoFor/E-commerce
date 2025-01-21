package rpc

import (
	"fmt"
	"github.com/MyGoFor/E-commerce/app/cart/utils"
	"github.com/MyGoFor/E-commerce/app/product/conf"
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
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
	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	if ProductClient == nil {
		fmt.Println("哈哈哈哈哈哈啊哈好", err)
	}
	utils.MustHandleError(err)
}
