package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/siralfbaez/mia-agentic-data-nervous-system-kfg-v3/pkg/observability"
	"github.com/siralfbaez/mia-agentic-data-nervous-system-kfg-v3/services/ai-agent/internal/logic"
)

func main() {
	// 1. Setup Context for Graceful Shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// 2. Initialize OpenTelemetry (Matching your KFG-v3 pattern)
	shutdown, err := observability.InitTelemetry(ctx, "ai-agent")
	if err != nil {
		log.Fatalf("failed to initialize telemetry: %v", err)
	}

	// Ensure telemetry flushes before the process exits
	defer func() {
		if err := shutdown(context.Background()); err != nil {
			log.Printf("error shutting down telemetry: %v", err)
		}
	}()

	log.Println("AI Agent (KFG-v3) starting up...")

	// 3. Initialize and Start your Agent Logic
	// Assuming your logic/agent.go has a function like Start(ctx)
	// If the function name is different in agent.go, update it here!
	err = logic.StartAgent(ctx)
	if err != nil {
		log.Fatalf("AI Agent failed to start: %v", err)
	}

	// 4. Wait for termination signal
	<-ctx.Done()
	log.Println("Shutting down AI Agent gracefully...")
}
