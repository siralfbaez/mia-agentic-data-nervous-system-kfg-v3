package observability

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
)

// ShutdownFunc is a cleanup function to be called on application exit
type ShutdownFunc func(context.Context) error

// InitTelemetry sets up both Tracing and Metrics for a service
func InitTelemetry(ctx context.Context, serviceName string) (ShutdownFunc, error) {
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String(serviceName),
			semconv.ServiceVersionKey.String("v3.0.0"),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	// 1. Setup Prometheus Metrics Exporter
	metricExporter, err := prometheus.New()
	if err != nil {
		return nil, fmt.Errorf("failed to create prometheus exporter: %w", err)
	}

	meterProvider := metric.NewMeterProvider(
		metric.WithResource(res),
		metric.WithReader(metricExporter),
	)
	otel.SetMeterProvider(meterProvider)

	// 2. Setup OTLP Trace Exporter (gRPC)
	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(traceExporter),
		sdktrace.WithResource(res),
	)
	otel.SetTracerProvider(tracerProvider)

	// Return a combined shutdown function
	return func(ctx context.Context) error {
		err1 := tracerProvider.Shutdown(ctx)
		err2 := meterProvider.Shutdown(ctx)
		if err1 != nil {
			return err1
		}
		return err2
	}, nil
}
