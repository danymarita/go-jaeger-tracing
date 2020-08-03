package main

import (
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
)

func startTrace(tracer opentracing.Tracer) {
	span := tracer.StartSpan("say-hello")
	span.Finish()
}

func parentSpan(tracer opentracing.Tracer) {
	parentSpan := tracer.StartSpan("parent-span")
	defer parentSpan.Finish()

	childSpan(tracer, parentSpan)
}

func childSpan(tracer opentracing.Tracer, parentSpan opentracing.Span) {
	childSpan := tracer.StartSpan(
		"child-span",
		opentracing.ChildOf(parentSpan.Context()),
	)
	defer childSpan.Finish()
}

func main() {
	// Sample configuration for testing. Use constant sampling to sample every trace
	// and enable LogSpan to log every span via configured Logger.
	cfg := jaegercfg.Configuration{
		ServiceName: "go-jaeger-tracing",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
			// Default Jaeger is localhost:16686, use LocalAgentHostPort property to define custome jaeger agent
			// LocalAgentHostPort: "HOST:PORT",
		},
	}

	// Example logger and metrics factory. Use github.com/uber/jaeger-client-go/log
	// and github.com/uber/jaeger-lib/metrics respectively to bind to real logging and metrics frameworks.
	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	// Initialize tracer with a logger and a metrics factory
	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)

	if err != nil {
		fmt.Printf("Error initiate tracer. Error : %v", err)
	}

	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

	startTrace(tracer)
	parentSpan(tracer)
}
