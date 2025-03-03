package service

import (
	"context"
	"fmt"
	"github.com/MyGoFor/E-commerce/app/frontend/infra/rpc"
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/eino"
	"log"

	ai "github.com/MyGoFor/E-commerce/app/frontend/hertz_gen/frontend/ai"
	"github.com/cloudwego/hertz/pkg/app"
)

type ToCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewToCartService(Context context.Context, RequestContext *app.RequestContext) *ToCartService {
	return &ToCartService{RequestContext: RequestContext, Context: Context}
}

func (h *ToCartService) Run(req *ai.AIReq) (resp *ai.Empty, err error) {
	content := req.Content
	fmt.Println("content:", content)

	uid := h.Context.Value("user_id").(int32)
	fmt.Println("uid:", uid)

	_, err = rpc.AIClient.PlaceOrder(h.Context, &eino.PlaceOrderReq{Uid: uid, Question: content})
	if err != nil {
		log.Println(err)
	}
	return
}
