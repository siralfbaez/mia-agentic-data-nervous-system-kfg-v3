package main

import (
	"context"
	"log"
	"net/http"

	"github.com/siralfbaez/mia-kfg-v3/pkg/observability"
	"github.com/siralfbaez/mia-kfg-v3/services/translation-engine/internal/mapping/cost_optimizer"
)

func main() {
	ctx := context.Background()
	
	// 1. Setup Observability
	shutdown, _ := observability.InitTelemetry(ctx, "translation-engine")
	defer shutdown(ctx)

	// 2. Initialize the Cost Optimizer Synapse
	optimizer := cost_optimizer.NewStreamOptimizer(0.7, 0.3)
	
	log.Println("KFG-v3 Translation Engine: Online & Optimizing...")
	
	// Logic to listen to Pub/Sub and run optimizer...
	http.ListenAndServe(":8081", nil)
}
