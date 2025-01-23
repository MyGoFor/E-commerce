package service

import (
	"context"
	checkout "github.com/MyGoFor/E-commerce/app/frontend/hertz_gen/frontend/checkout"
	"github.com/MyGoFor/E-commerce/app/frontend/infra/rpc"
	rpccheckout "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/checkout"
	rpcpayment "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/payment"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"log"
)

type CheckoutWaitingService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutWaitingService(Context context.Context, RequestContext *app.RequestContext) *CheckoutWaitingService {
	return &CheckoutWaitingService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutWaitingService) Run(req *checkout.CheckoutReq) (resp map[string]any, err error) {
	userId := uint32(h.Context.Value("user_id").(int32))

	_, err = rpc.CheckoutClient.Checkout(h.Context, &rpccheckout.CheckoutReq{
		UserId:    userId,
		Email:     req.Email,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Address: &rpccheckout.Address{
			Country:       req.Country,
			ZipCode:       req.Zipcode,
			City:          req.City,
			State:         req.Province,
			StreetAddress: req.Street,
		},
		CreditCard: &rpcpayment.CreditCardInfo{
			CreditCardNumber:          req.CardNum,
			CreditCardExpirationYear:  req.ExpirationYear,
			CreditCardExpirationMonth: req.ExpirationMonth,
			CreditCardCvv:             req.Cvv,
		},
	})
	if err != nil {
		log.Println("err :", err.Error())
		return nil, err
	}

	return utils.H{
		"Title":    "waiting",
		"redirect": "/checkout/result",
	}, nil
}
