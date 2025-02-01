package rpc

import (
	"github.com/MyGoFor/E-commerce/app/frontend/conf"
	frontendutils "github.com/MyGoFor/E-commerce/app/frontend/utils"
	"github.com/MyGoFor/E-commerce/common/clientsuite"
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/order/orderservice"
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"sync"
)

var (
	ProductClient  productcatalogservice.Client
	UserClient     userservice.Client
	CartClient     cartservice.Client
	CheckoutClient checkoutservice.Client
	OrderClient    orderservice.Client
	once           sync.Once
	ServiceName    = frontendutils.ServiceName
	RegistryAddr   = conf.GetConf().Hertz.RegistryAddr
	err            error
)

func Init() {
	once.Do(func() {
		initUserClient()
		initProductClient()
		initCartClient()
		initOrderClient()
		initCheckoutClient()
	})
}

func initUserClient() {
	UserClient, err = userservice.NewClient("user", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddress:    RegistryAddr,
	}))
	if err != nil {
		hlog.Fatal(err)
	}
}

func initProductClient() {
	ProductClient, err = productcatalogservice.NewClient("product", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddress:    RegistryAddr,
	}))
	if err != nil {
		hlog.Fatal(err)
	}
}

func initCartClient() {
	CartClient, err = cartservice.NewClient("cart", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddress:    RegistryAddr,
	}))
	if err != nil {
		hlog.Fatal(err)
	}
}

func initOrderClient() {
	OrderClient, err = orderservice.NewClient("order", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddress:    RegistryAddr,
	}))
	if err != nil {
		hlog.Fatal(err)
	}
}

func initCheckoutClient() {
	CheckoutClient, err = checkoutservice.NewClient("checkout", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddress:    RegistryAddr,
	}))
	if err != nil {
		hlog.Fatal(err)
	}
}
