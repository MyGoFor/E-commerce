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
	"errors"
	"fmt"
	myRedis "github.com/MyGoFor/E-commerce/app/cart/biz/dal/redis"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

type Cart struct {
	Base
	UserId    uint32 `json:"user_id"`
	ProductId uint32 `json:"product_id"`
	Qty       uint32 `json:"qty"`
}

func (c Cart) TableName() string {
	return "cart"
}

type CartQuery struct {
	ctx context.Context
	db  *gorm.DB
}

type CachedCartQuery struct {
	cartQuery   CartQuery
	cacheClient *redis.Client
	prefix      string
}

func (c CachedCartQuery) GetCartByUserId(userId uint32) (cartList []*Cart, err error) {
	cachedKey := fmt.Sprintf("%s_%s_%d", c.prefix, "cart_by_userId", userId)
	cachedResult := c.cacheClient.Get(c.cartQuery.ctx, cachedKey)

	err = func() error {
		err1 := cachedResult.Err()
		if err1 != nil {
			return err1
		}
		cachedResultByte, err2 := cachedResult.Bytes()
		if err2 != nil {
			return err2
		}
		err3 := json.Unmarshal(cachedResultByte, &cartList)
		if err3 != nil {
			return err3
		}
		return nil
	}()

	if err != nil {
		cartList, err = c.cartQuery.GetCartByUserId(userId)
		if err != nil {
			return nil, err
		}
		encoded, err := json.Marshal(cartList)
		if err != nil {
			return cartList, nil
		}
		_ = c.cacheClient.Set(c.cartQuery.ctx, cachedKey, encoded, 12*time.Hour)
	}
	return
}

func (c CartQuery) GetCartByUserId(userId uint32) (cartList []*Cart, err error) {
	err = c.db.WithContext(c.ctx).Model(&Cart{}).Where(&Cart{UserId: userId}).Find(&cartList).Error
	return
}

func NewCartQuery(ctx context.Context, db *gorm.DB) CartQuery {
	return CartQuery{
		ctx: ctx,
		db:  db,
	}
}

func NewCachedCartQuery(cq CartQuery, cacheClient *redis.Client) CachedCartQuery {
	return CachedCartQuery{
		cartQuery:   cq,
		cacheClient: cacheClient,
		prefix:      "E-commerce_shop",
	}
}

// 添加商品的时候应该更新缓存
func AddCart(db *gorm.DB, ctx context.Context, c *Cart) error {
	var find Cart
	err := db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserId: c.UserId, ProductId: c.ProductId}).First(&find).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if find.ID != 0 {
		err = db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserId: c.UserId, ProductId: c.ProductId}).UpdateColumn("qty", gorm.Expr("qty+?", c.Qty)).Error
	} else {
		err = db.WithContext(ctx).Model(&Cart{}).Create(c).Error
	}

	if err == nil {
		// 更新缓存
		cachedKey := fmt.Sprintf("E-commerce_shop_cart_by_userId_%d", c.UserId)
		var cartList []*Cart
		db.WithContext(ctx).Model(&Cart{}).Where("user_id = ?", c.UserId).Find(&cartList)
		encoded, _ := json.Marshal(cartList)
		myRedis.RedisClient.Set(ctx, cachedKey, encoded, 12*time.Hour)
	}

	return err
}

// 还要清理缓存
func EmptyCart(db *gorm.DB, ctx context.Context, userId uint32) error {
	if userId == 0 {
		return errors.New("user_is is required")
	}

	err := db.WithContext(ctx).Delete(&Cart{}, "user_id = ?", userId).Error
	if err == nil {
		// 清理缓存
		cachedKey := fmt.Sprintf("E-commerce_shop_cart_by_userId_%d", userId)
		myRedis.RedisClient.Del(ctx, cachedKey)
	}
	return err
}
