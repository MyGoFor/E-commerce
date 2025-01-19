package service

import (
	"context"
	"github.com/MyGoFor/E-commerce/app/product/biz/dal/mysql"
	"github.com/MyGoFor/E-commerce/app/product/module"
	product "github.com/MyGoFor/E-commerce/rpc_gen/product"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// Finish your business logic.
	var products []module.Product
	err = mysql.DB.Where("categories like ?", "%"+req.CategoryName+"%").Find(&products).Error
	if err != nil {
		return nil, err
	}
	resp = &product.ListProductsResp{
		Products: []*product.Product{},
	}
	for _, v := range products {
		p := &product.Product{
			Id:          v.Id,
			Name:        v.Name,
			Description: v.Description,
			Picture:     v.Picture,
			Price:       v.Price,
			Categories:  v.Categories,
		}
		resp.Products = append(resp.Products, p)
	}
	return resp, nil
}
