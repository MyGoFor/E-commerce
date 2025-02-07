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
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func (u User) TableName() string {
	return "user"
}

type UserQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func NewUserQuery(ctx context.Context, db *gorm.DB) UserQuery {
	return UserQuery{
		ctx: ctx,
		db:  db,
	}
}

func (u UserQuery) GetByEmail(email string) (user *User, err error) {
	err = u.db.WithContext(u.ctx).Model(&User{}).Where(&User{Email: email}).First(&user).Error
	return
}

type CacheUserQuery struct {
	userQuery   UserQuery
	cacheClient *redis.Client
	prefix      string
}

func NewCacheUserQuery(userQuery UserQuery, cacheClient *redis.Client) CacheUserQuery {
	return CacheUserQuery{
		userQuery:   userQuery,
		cacheClient: cacheClient,
	}
}

func (c CacheUserQuery) GetByEmail(email string) (user *User, err error) {
	cacheKey := fmt.Sprintf("%s_%s_%s", c.prefix, "user_by_email", email)
	cacheResult := c.cacheClient.Get(c.userQuery.ctx, cacheKey)
	err = func() error {
		err1 := cacheResult.Err()
		if err1 != nil {
			return err1
		}
		cachedResultByte, err2 := cacheResult.Bytes()
		if err2 != nil {
			return err2
		}
		err3 := json.Unmarshal(cachedResultByte, &user)
		if err3 != nil {
			return err3
		}
		return nil
	}()

	if err != nil {
		user, err = c.userQuery.GetByEmail(email)
		if err != nil {
			return &User{}, err
		}
		encoded, err := json.Marshal(user)
		if err != nil {
			return user, nil
		}
		_ = c.cacheClient.Set(c.userQuery.ctx, cacheKey, encoded, 12*time.Hour)
	}
	return
}

func Create(db *gorm.DB, ctx context.Context, user *User) error {
	return db.WithContext(ctx).Create(user).Error
}
