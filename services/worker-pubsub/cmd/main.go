package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/siralfbaez/mia-agentic-data-nervous-system-kfg-v3/pkg/observability"
	"github.com/siralfbaez/mia-agentic-data-nervous-system-kfg-v3/services/worker-pubsub/internal/processor"
)

func main() {
	// 1. Setup Context for Graceful Shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// 2. Initialize OpenTelemetry
	shutdown, err := observability.InitTelemetry(ctx, "worker-pubsub")
	if err != nil {
		log.Fatalf("failed to initialize telemetry: %v", err)
	}
	defer func() {
		if err := shutdown(context.Background()); err != nil {
			log.Printf("error shutting down telemetry: %v", err)
		}
	}()

	log.Println("Worker-PubSub (KFG-v3): Starting Nervous System Reflex Loop Engine...")

	// 3. Start the Processor Engine
	// NOTE: Ensure your engine.go has a function named StartEngine (or rename it below)
	if err := processor.StartEngine(ctx); err != nil {
		log.Fatalf("Worker engine failed: %v", err)
	}

	<-ctx.Done()
	log.Println("Shutting down Worker-PubSub Engine gracefully...")
}
