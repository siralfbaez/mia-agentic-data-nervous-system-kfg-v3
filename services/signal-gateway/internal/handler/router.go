package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/siralfbaez/mia-kfg-v3/pkg/observability"
	"github.com/siralfbaez/mia-kfg-v3/pkg/resilience"
	"go.opentelemetry.io/otel"
)

type Signal struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

func (h *IngestHandler) HandleSignal(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracer.Start(r.Context(), "Gateway-Ingress")
	defer span.End()

	var sig Signal
	if err := json.NewDecoder(r.Body).Decode(&sig); err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	// 1. Validation Synapse: Check if signal type is registered
	if sig.Type == "" {
		http.Error(w, "Missing Signal Type", http.StatusUnprocessableEntity)
		return
	}

	// 2. Publish to Pub/Sub (Async)
	// Staff Tip: Use a buffered channel or worker pool for extreme throughput
	go h.dispatchToStream(context.Background(), sig)

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"status": "queued", "id": span.SpanContext().TraceID().String()})
}
