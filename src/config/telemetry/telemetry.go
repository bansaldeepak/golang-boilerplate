package telemetry

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"
)

func SetupTelemetry(sugar *zap.SugaredLogger) (*trace.TracerProvider, error) {
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		sugar.Fatal(err)
		return nil, err
	}
	tp := trace.NewTracerProvider(trace.WithBatcher(exporter))
	otel.SetTracerProvider(tp)
	return tp, nil
}

func ShutdownTelemetry(tp *trace.TracerProvider) {
	_ = tp.Shutdown(context.Background())
}
