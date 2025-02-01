// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rpc

import (
	"github.com/MyGoFor/E-commerce/app/checkout/conf"
	"github.com/MyGoFor/E-commerce/common/clientsuite"
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/order/orderservice"
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"sync"

	"github.com/cloudwego/kitex/client"
)

var (
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
	OrderClient   orderservice.Client
	once          sync.Once
	err           error
	ServiceName   = conf.GetConf().Kitex.Service
	RegistryAddr  = conf.GetConf().Registry.RegistryAddress[0]
)

func InitClient() {
	once.Do(func() {
		initCartClient()
		initProductClient()
		initPaymentClient()
		initOrderClient()
	})
}

func initProductClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddress:    RegistryAddr,
		}),
	}
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	if err != nil {
		hlog.Fatal(err)
	}
}

func initCartClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddress:    RegistryAddr,
		}),
	}
	CartClient, err = cartservice.NewClient("cart", opts...)
	if err != nil {
		hlog.Fatal(err)
	}
}

func initPaymentClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddress:    RegistryAddr,
		}),
	}
	PaymentClient, err = paymentservice.NewClient("payment", opts...)
	if err != nil {
		hlog.Fatal(err)
	}
}

func initOrderClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddress:    RegistryAddr,
		}),
	}
	OrderClient, err = orderservice.NewClient("order", opts...)
	if err != nil {
		hlog.Fatal(err)
	}
}
