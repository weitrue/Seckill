/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2023/1/3 7:45 PM
 * Description:
 **/

package xtrace

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"

	"github.com/weitrue/Seckill/pkg/logger/meta"
	"github.com/weitrue/Seckill/pkg/logger/xlog"
)

// SetGlobalTracer ...
func SetGlobalTracer(ctx context.Context, tp trace.TracerProvider) {
	otel.SetTracerProvider(tp)
	xlog.WithContext(ctx).Info("set global tracer", meta.NewField("trace", tp))
}

type options struct {
	propagator propagation.TextMapPropagator
}

// Option is tracing option.
type Option func(*options)

// Tracer is otel span tracer
type Tracer struct {
	tracer trace.Tracer
	kind   trace.SpanKind
	opt    *options
}

// NewTracer create tracer instance
func NewTracer(kind trace.SpanKind, opts ...Option) *Tracer {
	op := options{
		// propagator: propagation.NewCompositeTextMapPropagator(Metadata{}, propagation.Baggage{}, propagation.TraceContext{}),
		propagator: propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}),
	}
	for _, o := range opts {
		o(&op)
	}
	return &Tracer{tracer: otel.Tracer("tracespace"), kind: kind, opt: &op}
}

// Start start tracing span
func (t *Tracer) Start(ctx context.Context, operation string, carrier propagation.TextMapCarrier, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	if (t.kind == trace.SpanKindServer || t.kind == trace.SpanKindConsumer) && carrier != nil {
		ctx = t.opt.propagator.Extract(ctx, carrier)
	}
	opts = append(opts, trace.WithSpanKind(t.kind))

	ctx, span := t.tracer.Start(ctx, operation, opts...)

	if (t.kind == trace.SpanKindClient || t.kind == trace.SpanKindProducer) && carrier != nil {
		t.opt.propagator.Inject(ctx, carrier)
	}
	return ctx, span
}
