package service

import (
	"context"
	"github.com/MyGoFor/E-commerce/app/product/biz/dal/mysql"
	"github.com/MyGoFor/E-commerce/app/product/module"
	product "github.com/MyGoFor/E-commerce/rpc_gen/product"
)

type AddProductService struct {
	ctx context.Context
} // NewAddProductService new AddProductService
func NewAddProductService(ctx context.Context) *AddProductService {
	return &AddProductService{ctx: ctx}
}

// Run create note info
func (s *AddProductService) Run(req *product.AddProductReq) (resp *product.AddProductResp, err error) {
	// Finish your business logic.
	p := module.Product{
		Name:        req.Product.Name,
		Description: req.Product.Description,
		Picture:     req.Product.Picture,
		Price:       req.Product.Price,
		Categories:  req.Product.Categories,
	}
	err = mysql.DB.Create(&p).Error
	//err = redis.RedisClient.HSet(context.Background(), strconv.Itoa(int(p.Id)), "").Err()
	if err != nil {
		return nil, err
	}
	return &product.AddProductResp{Id: p.Id}, nil
}
