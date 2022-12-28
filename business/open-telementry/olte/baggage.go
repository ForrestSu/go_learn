package olte

import (
	"context"

	"go.opentelemetry.io/otel/baggage"
)

// AddBaggage add baggage
func AddBaggage(ctx context.Context) context.Context {
	member, _ := baggage.NewMember("userid", "Bob")
	bag, _ := baggage.New(member)
	return baggage.ContextWithBaggage(ctx, bag)
}
