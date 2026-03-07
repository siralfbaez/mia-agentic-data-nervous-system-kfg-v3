package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"go.opentelemetry.io/otel"
)

// Define tracer ONCE for the whole package here
var tracer = otel.Tracer("signal-gateway")

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

	if sig.Type == "" {
		http.Error(w, "Missing Signal Type", http.StatusUnprocessableEntity)
		return
	}

	sigBytes, _ := json.Marshal(sig)
	// Passing ctx to satisfy the 'unused' error and the receiver method
	go h.dispatchToStream(ctx, sigBytes)

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{
		"status":   "queued",
		"trace_id": span.SpanContext().TraceID().String(),
	})
}

func (h *IngestHandler) dispatchToStream(ctx context.Context, data []byte) {
	_, span := tracer.Start(ctx, "Dispatch-To-KFG")
	defer span.End()
	_ = data
}