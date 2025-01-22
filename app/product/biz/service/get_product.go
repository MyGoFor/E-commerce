package service

import (
	"context"
	"github.com/MyGoFor/E-commerce/app/product/biz/dal/mysql"
	"github.com/MyGoFor/E-commerce/app/product/module"
	product "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/product"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	// Finish your business logic.
	p := module.Product{}
	err = mysql.DB.Where("id = ?", req.Id).First(&p).Error
	if err != nil {
		return nil, err
	}
	return &product.GetProductResp{Product: &product.Product{
		Id:          p.Id,
		Name:        p.Name,
		Description: p.Description,
		Picture:     p.Picture,
		Price:       p.Price,
		Categories:  p.Categories,
	}}, nil
}
