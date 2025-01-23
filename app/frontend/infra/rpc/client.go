package rpc

import (
	"github.com/MyGoFor/E-commerce/app/frontend/conf"
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/order/orderservice"
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"sync"
)

var (
	ProductClient  productcatalogservice.Client
	UserClient     userservice.Client
	CartClient     cartservice.Client
	CheckoutClient checkoutservice.Client
	OrderClient    orderservice.Client
	once           sync.Once
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

func initCartClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	if err != nil {
		hlog.Fatal(err)
	}
	CartClient, err = cartservice.NewClient("cart", client.WithResolver(r))
	if err != nil {
		hlog.Fatal(err)
	}
}

func initOrderClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	if err != nil {
		hlog.Fatal(err)
	}
	OrderClient, err = orderservice.NewClient("order", client.WithResolver(r))
	if err != nil {
		hlog.Fatal(err)
	}
}

func initCheckoutClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	if err != nil {
		hlog.Fatal(err)
	}
	CheckoutClient, err = checkoutservice.NewClient("checkout", client.WithResolver(r))
	if err != nil {
		hlog.Fatal(err)
	}
}
