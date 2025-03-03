package AiModel

import (
	"context"
	"github.com/MyGoFor/E-commerce/app/eino/infra/rpc"
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/cart"
	_ "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/order"
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
	"log"
	"strings"
)

func PlaceModel(ctx context.Context, question string, uid int32) error {

	// 初始化模型
	model, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: "697cee32-664b-4102-bec5-fc859b1d5158",
		Region: "cn-beijing",
		Model:  "ep-20250129165433-s8xk8",
	})
	if err != nil {
		panic(err)
	}

	// 创建模板，使用 FString 格式
	template := prompt.FromMessages(schema.FString,
		// 系统消息模板
		schema.SystemMessage("你是一个{role}。你需要用{style}的语气回答问题。你的目标整理用户列出的订单里面商品并以一个空格分隔列出每个商品名字"),

		// 插入需要的对话历史（新对话的话这里不填）
		schema.MessagesPlaceholder("chat_history", true),

		// 用户消息模板
		schema.UserMessage("问题: {question}"),
	)

	// 使用模板生成消息
	messages, err := template.Format(context.Background(), map[string]any{
		"role":     "服务员",
		"style":    "严谨",
		"question": question,
		//对话历史（这个例子里模拟两轮对话历史）
		"chat_history": []*schema.Message{
			schema.UserMessage("你好"),
			schema.AssistantMessage("嘿！我是你的自动下单助手！请列出你的下单商品名称以及数量", nil),
			schema.UserMessage("我买了两个麻衣学姐手办以及一个帅哥自拍"),
			schema.AssistantMessage("麻衣学姐手办 麻衣学姐手办 帅哥自拍", nil),
		},
	})

	// 生成回复
	response, err := model.Generate(ctx, messages)
	if err != nil {
		panic(err)
	}
	log.Println(response.Content)

	slice := strings.Split(response.Content, " ")
	for _, s := range slice {
		log.Println(s)
		ProductResp, err := rpc.ProductClient.SearchProducts(ctx, &product.SearchProductsReq{Query: s})
		if err != nil {
			return err
		}
		item := ProductResp.Results[0]
		_, err = rpc.CartClient.AddItem(ctx, &cart.AddItemReq{
			UserId: uint32(uid),
			Item: &cart.CartItem{
				ProductId: item.Id,
				Quantity:  1,
			},
		})
		if err != nil {
			return err
		}
	}
	return nil
	////rpc创建order
	//slice := strings.Split(response.Content, " ")
	//var orderItems []*order.OrderItem
	//for _, s := range slice {
	//	log.Println(s)
	//	resp, err := rpc.ProductClient.SearchProducts(ctx, &product.SearchProductsReq{Query: s})
	//	if err != nil {
	//		return nil, err
	//	}
	//	item := resp.Results[0]
	//	o := &order.OrderItem{
	//		Item: &cart.CartItem{
	//			ProductId: item.Id,
	//			Quantity:  1,
	//		},
	//		Cost: item.Price,
	//	}
	//	orderItems = append(orderItems, o)
	//}
	//resp, err := rpc.OrderClient.PlaceOrder(ctx, &order.PlaceOrderReq{
	//	UserId:       0,
	//	UserCurrency: "",
	//	Address:      nil,
	//	Email:        "",
	//	OrderItems:   orderItems,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//
	//return resp, nil

}
