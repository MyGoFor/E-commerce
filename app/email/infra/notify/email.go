package notify

import (
	"github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/email"
	"github.com/kr/pretty"
)

type NoopEmail struct{}

func (e *NoopEmail) Send(req *email.EmailReq) error {
	// 直接打印，不做真实邮件发送操作
	pretty.Printf("%v\n", req)
	return nil
}

func NewNoopEmail() NoopEmail {
	return NoopEmail{}
}
