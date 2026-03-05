package resilience

import (
	"context"
	"errors"
	"time"

	"github.com/sony/gobreaker" // A standard Staff-level choice for Go
	"github.com/siralfbaez/mia-kfg-v3/pkg/observability"
	"go.opentelemetry.io/otel/attribute"
)

var (
	ErrCircuitOpen = errors.New("circuit breaker is open - downstream service unhealthy")
)

// Policy wraps the circuit breaker and retry logic
type Policy struct {
	cb *gobreaker.CircuitBreaker
}

func NewResiliencePolicy(name string) *Policy {
	settings := gobreaker.Settings{
		Name:        name,
		MaxRequests: 5,
		Interval:    30 * time.Second,
		Timeout:     10 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return counts.Requests >= 3 && failureRatio >= 0.6
		},
	}

	return &Policy{
		cb: gobreaker.NewCircuitBreaker(settings),
	}
}

// Execute wraps a function call with circuit breaking logic
func (p *Policy) Execute(ctx context.Context, operation func() (interface{}, error)) (interface{}, error) {
	result, err := p.cb.Execute(operation)
	
	if err != nil {
		if err == gobreaker.ErrOpenState {
			return nil, ErrCircuitOpen
		}
		return nil, err
	}

	return result, nil
}
