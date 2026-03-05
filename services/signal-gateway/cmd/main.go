package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/siralfbaez/mia-kfg-v3/pkg/observability"
	"github.com/siralfbaez/mia-kfg-v3/services/signal-gateway/internal/handler"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Initialize OpenTelemetry Wrapper
	shutdown, err := observability.InitTelemetry(ctx, "signal-gateway")
	if err != nil {
		log.Fatalf("failed to initialize telemetry: %v", err)
	}
	defer shutdown(context.Background())

	h := &handler.IngestHandler{
		topic: "signals.raw.v3",
	}

	server := &http.Server{
		Addr:    ":8080",
		Handler: h,
	}

	log.Println("Signal Gateway (KFG-v3) listening on :8080")
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down Gateway gracefully...")
}
