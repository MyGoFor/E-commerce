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

package model

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/redis/go-redis/v9"
	"time"

	"gorm.io/gorm"
)

type Consignee struct {
	Email string

	StreetAddress string
	City          string
	State         string
	Country       string
	ZipCode       int32
}

type OrderState string

const (
	OrderStatePlaced   OrderState = "placed"
	OrderStatePaid     OrderState = "paid"
	OrderStateCanceled OrderState = "canceled"
)

type Order struct {
	Base
	OrderId      string `gorm:"uniqueIndex;size:256"`
	UserId       uint32
	UserCurrency string
	Consignee    Consignee   `gorm:"embedded"`
	OrderItems   []OrderItem `gorm:"foreignKey:OrderIdRefer;references:OrderId"`
	OrderState   OrderState
}

type OrderQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func NewOrderQuery(ctx context.Context, db *gorm.DB) OrderQuery {
	return OrderQuery{
		ctx: ctx,
		db:  db,
	}
}

func (o OrderQuery) GetOrder(userId uint32, orderId string) (order Order, err error) {
	err = o.db.WithContext(o.ctx).Model(&Order{}).Where(&Order{UserId: userId, OrderId: orderId}).First(&order).Error
	return
}

type CachedOrderQuery struct {
	orderQuery  OrderQuery
	cacheClient *redis.Client
	prefix      string
}

func NewCachedOrderQuery(oq OrderQuery, cacheClient *redis.Client) CachedOrderQuery {
	return CachedOrderQuery{
		orderQuery:  oq,
		cacheClient: cacheClient,
		prefix:      "E-commerce_shop",
	}
}

func (o CachedOrderQuery) GetOrder(userId uint32, orderId string) (order Order, err error) {
	cacheKey := fmt.Sprintf("%s_%s_%d_%s", o.prefix, "order_by_id", userId, orderId)
	cacheResult := o.cacheClient.Get(o.orderQuery.ctx, cacheKey)
	err = func() error {
		err1 := cacheResult.Err()
		if err1 != nil {
			return err1
		}
		cachedResultByte, err2 := cacheResult.Bytes()
		if err2 != nil {
			return err2
		}
		err3 := json.Unmarshal(cachedResultByte, &order)
		if err3 != nil {
			return err3
		}
		return nil
	}()

	if err != nil {
		order, err = o.orderQuery.GetOrder(userId, orderId)
		if err != nil {
			return Order{}, err
		}
		encoded, err := json.Marshal(order)
		if err != nil {
			return order, nil
		}
		_ = o.cacheClient.Set(o.orderQuery.ctx, cacheKey, encoded, 12*time.Hour)
	}
	return
}

func (o Order) TableName() string {
	return "order"
}

func ListOrder(db *gorm.DB, ctx context.Context, userId uint32) (orders []Order, err error) {
	err = db.Model(&Order{}).Where(&Order{UserId: userId}).Preload("OrderItems").Find(&orders).Error
	return
}

func UpdateOrderState(db *gorm.DB, ctx context.Context, userId uint32, orderId string, state OrderState) error {
	return db.Model(&Order{}).Where(&Order{UserId: userId, OrderId: orderId}).Update("order_state", state).Error
}
