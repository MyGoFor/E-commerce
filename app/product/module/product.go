package module

import (
	"context"
	"database/sql/driver"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

type stringArray []string

func (sa stringArray) Value() (driver.Value, error) {
	return json.Marshal(sa)
}

func (sa *stringArray) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan StringArray: %v", value)
	}
	return json.Unmarshal(bytes, sa)
}

type Product struct {
	Id          uint32      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string      `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Picture     string      `protobuf:"bytes,4,opt,name=picture,proto3" json:"picture,omitempty"`
	Price       float32     `protobuf:"fixed32,5,opt,name=price,proto3" json:"price,omitempty"`
	Categories  stringArray `protobuf:"bytes,6,rep,name=categories,proto3" json:"categories,omitempty" gorm:"type:json"`
}

func (p Product) TableName() string {
	return "products"
}

// 数据库直接查询
type ProductQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func (p ProductQuery) GetById(productId uint32) (product Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).Where(&Product{Id: productId}).First(&product).Error
	return
}

func NewProductQuery(ctx context.Context, db *gorm.DB) ProductQuery {
	return ProductQuery{ctx: ctx, db: db}
}

// 增加缓存
type CachedProductQuery struct {
	productQuery ProductQuery
	cacheClient  *redis.Client
	prefix       string
}

func (c CachedProductQuery) GetByID(productId uint32) (product Product, err error) {
	cacheKey := fmt.Sprintf("%s_%s_%d", c.prefix, "product_by_id", productId)
	cacheResult := c.cacheClient.Get(c.productQuery.ctx, cacheKey)

	// 闭包构建一个错误链
	err = func() error {
		err1 := cacheResult.Err()
		if err1 != nil {
			return err1
		}
		cachedResultByte, err2 := cacheResult.Bytes()
		if err2 != nil {
			return err2
		}
		err3 := json.Unmarshal(cachedResultByte, &product)
		if err3 != nil {
			return err3
		}
		return nil
	}()

	if err != nil {
		// 数据库查询
		product, err = c.productQuery.GetById(productId)
		if err != nil {
			return Product{}, err
		}
		// 序列化之后放入缓存
		encoded, err := json.Marshal(product)
		if err != nil {
			return product, nil
		}
		_ = c.cacheClient.Set(c.productQuery.ctx, cacheKey, encoded, 12*time.Hour)
	}
	return
}

func NewCachedProductQuery(pq ProductQuery, cacheClient *redis.Client) CachedProductQuery {
	return CachedProductQuery{productQuery: pq, cacheClient: cacheClient, prefix: "E-commerce_shop"}
}
