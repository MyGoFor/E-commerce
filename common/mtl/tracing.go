package mtl

import (
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"os"
)

func InitTracing(serviceName string) provider.OtelProvider {
	p := provider.NewOpenTelemetryProvider(
		provider.WithExportEndpoint(os.Getenv("OTEL_EXPORTER_OTLP_TRACES_ENDPOINT")),
		provider.WithServiceName(serviceName),
		provider.WithInsecure(),
		provider.WithEnableMetrics(false), // 前面用了prometheus,这里关闭otel的metrics
	)
	return p
}
