package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"go.opentelemetry.io/otel" // Added for tracing
)

var tracer = otel.Tracer("signal-gateway")

type Signal struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

func (h *IngestHandler) HandleSignal(w http.ResponseWriter, r *http.Request) {
	// FIX 1: Uncomment and define ctx/span properly
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

	// FIX 2: Convert struct to bytes before dispatching,
	// and match the new function signature below
	sigBytes, _ := json.Marshal(sig)
	go h.dispatchToStream(context.Background(), sigBytes)

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{
		"status":   "queued",
		"trace_id": span.SpanContext().TraceID().String(),
	})
}

// FIX 3: Update signature to match how it's called (added context)
func (h *IngestHandler) dispatchToStream(ctx context.Context, data []byte) {
	_, span := tracer.Start(ctx, "Dispatch-To-KFG")
	defer span.End()

	// Placeholder for internal routing logic
	_ = data
}