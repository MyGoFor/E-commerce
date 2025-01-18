package service

import (
	"context"
	home "github.com/MyGoFor/E-commerce/app/frontend/hertz_gen/frontend/home"
	"github.com/cloudwego/hertz/pkg/app"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *home.Empty) (resp map[string]any, err error) {
	items := []map[string]any{
		{"Name": "02.0", "Price": 1.99, "Picture": "https://tuchuang.hch1212.online/img/02.webp"},
		{"Name": "02.1", "Price": 1.00, "Picture": "https://tuchuang.hch1212.online/img/021.webp"},
		{"Name": "02.3", "Price": 1.05, "Picture": "https://tuchuang.hch1212.online/img/023.webp"},
	}
	resp = map[string]any{
		"Items": items,
		"Title": "E-commerce",
	}
	return
}
