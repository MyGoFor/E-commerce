package rpc

import (
	"context"
	"github.com/MyGoFor/E-commerce/app/frontend/conf"
	frontendutils "github.com/MyGoFor/E-commerce/app/frontend/utils"
	"github.com/MyGoFor/E-commerce/common/clientsuite"
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/casbin/casbinservice"
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/order/orderservice"
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/product"
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/fallback"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"sync"
)

var (
	CasbinClient   casbinservice.Client
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
		initCasbinClient()
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
	// 添加熔断机制
	cbs := circuitbreak.NewCBSuite(func(ri rpcinfo.RPCInfo) string {
		return circuitbreak.RPCInfo2Key(ri)
	})
	cbs.UpdateServiceCBConfig("frontend/product/GetProduct",
		circuitbreak.CBConfig{Enable: true, ErrRate: 0.5, MinSample: 2})

	ProductClient, err = productcatalogservice.NewClient("product", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddress:    RegistryAddr,
	}), client.WithCircuitBreaker(cbs), client.WithFallback( // 添加 fallback
		fallback.NewFallbackPolicy(fallback.UnwrapHelper(func(ctx context.Context, req, resp interface{}, err error) (fbResp interface{}, fbErr error) {
			if err == nil {
				return resp, nil
			}
			methodName := rpcinfo.GetRPCInfo(ctx).To().Method()
			if methodName != "ListProducts" {
				return resp, err
			}
			return &product.ListProductsResp{
				Products: []*product.Product{
					{
						Picture:     "http://tuchuang.hch1212.online/img/0212.webp",
						Id:          12,
						Name:        "02.12",
						Description: "02.12",
						Price:       1.65,
					},
				},
			}, nil
		})),
	))
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

func initCasbinClient() {
	CasbinClient, err = casbinservice.NewClient("casbin", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddress:    RegistryAddr,
	}))
	if err != nil {
		hlog.Fatal(err)
	}
}
