package AiModel

import (
	"context"
	"github.com/MyGoFor/E-commerce/app/eino/infra/rpc"
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
	"strconv"
)

func SearchModel(question string) (*order.ListOrderResp, error) {
	ctx := context.Background()

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
		schema.SystemMessage("你是一个{role}。你需要用{style}的语气回答问题。你的目标是提取用户的话里面的用户id"),

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
			schema.UserMessage("我的用户id为1,请查询我的订单"),
			schema.AssistantMessage("1", nil),
		},
	})

	// 生成回复
	response, err := model.Generate(ctx, messages)
	if err != nil {
		panic(err)
	}
	uid, err := strconv.Atoi(response.String())
	resp, err := rpc.OrderClient.ListOrder(ctx, &order.ListOrderReq{UserId: uint32(uid)})
	if err != nil {
		return nil, err
	}
	return resp, nil

}
