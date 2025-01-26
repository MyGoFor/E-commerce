package serversuite

import (
	"github.com/MyGoFor/E-commerce/common/mtl"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
	"log"
)

type CommonServerSuite struct {
	CurrentServiceName string
	RegistryAddress    string
}

func (s CommonServerSuite) Options() []server.Option {
	opts := []server.Option{
		server.WithMetaHandler(transmeta.ClientHTTP2Handler),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: s.CurrentServiceName}),
		server.WithTracer(prometheus.NewServerTracer("", "", prometheus.WithDisableServer(true), prometheus.WithRegistry(mtl.Registry))),
		server.WithSuite(tracing.NewServerSuite()),
	}

	r, err := consul.NewConsulRegister(s.RegistryAddress)
	if err != nil {
		log.Fatal(err)
	}
	opts = append(opts, server.WithRegistry(r))

	return opts
}
