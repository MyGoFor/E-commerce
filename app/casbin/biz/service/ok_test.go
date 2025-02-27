package service

import (
	"context"
	casbin "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/casbin"
	"testing"
)

func TestOk_Run(t *testing.T) {
	ctx := context.Background()
	s := NewOkService(ctx)
	// init req and assert value

	req := &casbin.OkReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
