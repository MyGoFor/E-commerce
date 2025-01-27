package service

import (
	"context"
	upgrade "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/upgrade"
	"testing"
)

func TestOk_Run(t *testing.T) {
	ctx := context.Background()
	s := NewOkService(ctx)
	// init req and assert value

	req := &upgrade.OkReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
