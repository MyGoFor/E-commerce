package email

import (
	"github.com/MyGoFor/E-commerce/app/email/infra/mq"
	"github.com/MyGoFor/E-commerce/app/email/infra/notify"
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/email"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

func ConsumerInit() {
	// 订阅一个主题
	sub, err := mq.Nc.Subscribe("email", func(msg *nats.Msg) {
		var req email.EmailReq
		// 消息格式是protobuf
		err := proto.Unmarshal(msg.Data, &req)
		if err != nil {
			klog.Error("proto.Unmarshal", err)
			return
		}

		// 发送消息
		noopEmail := notify.NewNoopEmail()
		_ = noopEmail.Send(&req)
	})
	if err != nil {
		panic(err)
	}

	// 注册kitex的hook，在服务关闭时取消订阅
	server.RegisterShutdownHook(func() {
		_ = sub.Unsubscribe()
		mq.Nc.Close()
	})
}
