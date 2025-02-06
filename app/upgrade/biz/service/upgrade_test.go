package service

import (
	"context"
	upgrade "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/upgrade"
	"testing"
)

func TestUpgrade_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpgradeService(ctx)
	// init req and assert value

	req := &upgrade.UpgradeReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
