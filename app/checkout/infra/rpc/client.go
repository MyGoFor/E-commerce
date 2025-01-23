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
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/order/orderservice"
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	consul "github.com/kitex-contrib/registry-consul"
	"sync"

	"github.com/cloudwego/kitex/client"
)

var (
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
	OrderClient   orderservice.Client
	once          sync.Once
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
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		hlog.Fatal(err)
	}
	ProductClient, err = productcatalogservice.NewClient("product", client.WithResolver(r))
	if err != nil {
		hlog.Fatal(err)
	}
}

func initCartClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		hlog.Fatal(err)
	}
	CartClient, err = cartservice.NewClient("cart", client.WithResolver(r))
	if err != nil {
		hlog.Fatal(err)
	}
}

func initPaymentClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		hlog.Fatal(err)
	}
	PaymentClient, err = paymentservice.NewClient("payment", client.WithResolver(r))
	if err != nil {
		hlog.Fatal(err)
	}
}

func initOrderClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		hlog.Fatal(err)
	}
	OrderClient, err = orderservice.NewClient("order", client.WithResolver(r))
	if err != nil {
		hlog.Fatal(err)
	}
}
