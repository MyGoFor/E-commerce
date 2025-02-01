package mtl

import (
	"github.com/kitex-contrib/obs-opentelemetry/provider"
)

func InitTracing(serviceName string) provider.OtelProvider {
	p := provider.NewOpenTelemetryProvider(
		provider.WithExportEndpoint("127.0.0.1:4317"), // 本地用这行
		//provider.WithExportEndpoint(os.Getenv("OTEL_EXPORTER_OTLP_TRACES_ENDPOINT")), // docker容器部署用这行
		provider.WithServiceName(serviceName),
		provider.WithInsecure(),
		provider.WithEnableMetrics(false), // 前面用了prometheus,这里关闭otel的metrics
	)
	return p
}
