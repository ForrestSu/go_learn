package main

import (
	"context"
	"fmt"
	"os"

	"git.code.oa.com/tpstelemetry/tps-sdk-go/api"
	"git.woa.com/opentelemetry/opentelemetry-go-ecosystem/otelzap"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv"
	semconv "go.opentelemetry.io/otel/semconv/v1.8.0"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	tracer, err := initTracer()
	if err != nil {
		panic(err)
	}
}

func initTracer() (*sdktrace.TracerProvider, error) {
	// 标准输出
	exporterStdout, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		return nil, err
	}

	// 文件输出
	f, err := os.Create("traces_client.json")
	if err != nil {
		return nil, err
	}

	exporterFile, err := stdouttrace.New(
		stdouttrace.WithWriter(f),
		stdouttrace.WithPrettyPrint(),
	)
	if err != nil {
		return nil, err
	}

	// OTLP 输出到天机阁
	exporterTps, err := otlptracegrpc.New(context.Background(),
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint("9.134.111.27:12520"),
		otlptracegrpc.WithCompressor("gzip"),
		otlptracegrpc.WithHeaders(map[string]string{"X-Tps-TenantID": "default"}),
		otlptracegrpc.WithDialOption(grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(4194304))),
	)
	if err != nil {
		return nil, err
	}

	res, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			"", // semconv.SchemaURL,
			semconv.ServiceNameKey.String("client_demo"),
			semconv.ServiceVersionKey.String("v0.1.0"),
			attribute.String("environment", "test"),
			attribute.Key("tps.tenant.id").String("default"),
		),
	)

	// For the demonstration, use sdktrace.AlwaysSample sampler to sample all traces.
	// In a production application, use sdktrace.ProbabilitySampler with a desired probability.
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporterStdout),
		sdktrace.WithBatcher(exporterFile),
		sdktrace.WithBatcher(exporterTps),
		sdktrace.WithResource(res),
	)

	otel.SetTracerProvider(tp)

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return tp, nil
}

func initLogger() *zap.Logger {
	exp, err := otlplog.NewExporter(otlplog.WithInsecure(),
		otlplog.WithAddress("9.134.111.27:12520"))
	if err != nil {
		fmt.Println(err)
	}

	res := resource.NewWithAttributes(
		"", api.TpsTenantIDKey.String("default"),
	)

	core := otelzap.NewBatchCore(otelzap.NewBatchWriteSyncer(exp, res))
	return zap.New(core)
}
