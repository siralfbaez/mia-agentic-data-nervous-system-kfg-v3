package handler

import (
	"net/http"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

var (
	tracer = otel.Tracer("signal-gateway")
	meter  = otel.Meter("signal-gateway")
)

type IngestHandler struct {
	pubSubClient interface{} // Simplified for snippet
	topic        string
}

func (h *IngestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 1. Start Trace Span
	ctx, span := tracer.Start(r.Context(), "IngestSignal",
		trace.WithAttributes(attribute.String("http.method", r.Method)),
	)
	defer span.End()

	// 2. Metric: Increment Request Count
	counter, _ := meter.Int64Counter("gateway.requests_total")
	counter.Add(ctx, 1)

	// 3. Schema Validation (Staff-level requirement)
	// We check if the incoming signal matches our Protobuf/Avro contracts
	if err := h.validateSchema(r); err != nil {
		span.RecordError(err)
		http.Error(w, "Invalid Schema Contract", http.StatusBadRequest)
		return
	}

	// 4. Async Dispatch to Pub/Sub
	start := time.Now()
	err := h.publishToStream(ctx, r.Body)

	// 5. Metric: Latency for Stream Write
	latency, _ := meter.Float64Histogram("gateway.publish_latency_ms")
	latency.Record(ctx, float64(time.Since(start).Milliseconds()))

	if err != nil {
		http.Error(w, "Internal Ingestion Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
