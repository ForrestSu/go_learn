package olte

import (
	"context"
	"encoding/json"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// AddEvent add event
func AddEvent(ctx context.Context, req interface{}, rsp interface{}) {
	if span := trace.SpanFromContext(ctx); span.SpanContext().IsSampled() {
		span.AddEvent("incoming", trace.WithAttributes(
			attribute.Key("req").String(ProtoMessageToJSONString(req)),
		))
		span.AddEvent("outgoing", trace.WithAttributes(
			attribute.Key("rsp").String(ProtoMessageToJSONString(rsp)),
		))
	}
}

// AddAttributes add trace attributes
func AddAttributes(ctx context.Context, attrs ...attribute.KeyValue) {
	if span := trace.SpanFromContext(ctx); span.SpanContext().IsSampled() {
		span.SetAttributes(attrs...)
	}
}

func ProtoMessageToJSONString(msg interface{}) string {
	s, _ := json.Marshal(msg)
	return string(s)
}

func test() {
	AddAttributes(nil, "")
	AddEvent()
}
