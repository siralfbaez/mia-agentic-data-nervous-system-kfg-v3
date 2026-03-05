#!/bin/bash
# MIA KFG-v3 Module Aligner

REPO_FULL="github.com/siralfbaez/mia-agentic-data-nervous-system-kfg-v3"

# 1. Fix Shared Packages
modules=("pkg/observability" "pkg/resilience" "pkg/encoding-utils")
for m in "${modules[@]}"; do
    cd $m && go mod edit -module $REPO_FULL/$m && go mod tidy && cd ../..
done

# 2. Fix Services with Local Replaces
services=("services/ai-agent" "services/signal-gateway" "services/translation-engine" "services/worker-pubsub" "services/contract-validator")
for s in "${services[@]}"; do
    cd $s
    go mod edit -module $REPO_FULL/$s
    # Point to local packages instead of the internet
    go mod edit -replace github.com/siralfbaez/mia-kfg-v3/pkg/observability=../../pkg/observability
    go mod edit -replace github.com/siralfbaez/mia-kfg-v3/pkg/resilience=../../pkg/resilience
    go mod tidy
    cd ../..
done

go work sync
