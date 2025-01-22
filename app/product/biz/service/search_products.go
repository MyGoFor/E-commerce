package service

import (
	"context"
	"github.com/MyGoFor/E-commerce/app/product/biz/dal/mysql"
	"github.com/MyGoFor/E-commerce/app/product/module"
	product "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/product"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	// Finish your business logic.
	var products []*module.Product
	err = mysql.DB.Where("name like ?", "%"+req.Query+"%").Find(&products).Error
	if err != nil {
		return nil, err
	}
	resp = &product.SearchProductsResp{Results: []*product.Product{}}
	for _, v := range products {
		p := &product.Product{
			Id:          v.Id,
			Name:        v.Name,
			Description: v.Description,
			Picture:     v.Picture,
			Price:       v.Price,
			Categories:  v.Categories,
		}
		resp.Results = append(resp.Results, p)
	}
	return resp, nil
}
