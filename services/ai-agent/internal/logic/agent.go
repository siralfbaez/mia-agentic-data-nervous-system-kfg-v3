package logic

import (
	"context"
	"errors"
	"fmt"

	"cloud.google.com/go/vertexai/genai"
	"github.com/siralfbaez/mia-agentic-data-nervous-system-kfg-v3/pkg/observability"
	"github.com/siralfbaez/mia-agentic-data-nervous-system-kfg-v3/pkg/resilience"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

var (
	tracer = otel.Tracer("ai-agent")
)

// AgentLogic orchestrates reasoning over the Knowledge Flow Graph (KFG)
type AgentLogic struct {
	client     *genai.Client
	model      *genai.GenerativeModel
	resilience *resilience.Policy // Protects the Nervous System from AI latency
}

// NewAgentLogic initializes the Gemini client and resilience policies
func NewAgentLogic(ctx context.Context, projectID, region string) (*AgentLogic, error) {
	client, err := genai.NewClient(ctx, projectID, region)
	if err != nil {
		return nil, fmt.Errorf("failed to init vertex ai: %w", err)
	}

	// Using 1.5 Pro for deep reasoning over KFG state
	model := client.GenerativeModel("gemini-1.5-pro")

	// Initialize Circuit Breaker Synapse
	policy := resilience.NewResiliencePolicy("gemini-reasoning-cb")

	return &AgentLogic{
		client:     client,
		model:      model,
		resilience: policy,
	}, nil
}

// ProcessEnrichedSignal takes a KFG state and returns an autonomous decision
func (a *AgentLogic) ProcessEnrichedSignal(ctx context.Context, kfgState string) (string, error) {
	ctx, span := tracer.Start(ctx, "AI-Reasoning-Task")
	defer span.End()

	span.SetAttributes(attribute.Int("state.length", len(kfgState)))

	// Execute via Resilience Policy to prevent cascading failures
	result, err := a.resilience.Execute(ctx, func() (interface{}, error) {
		prompt := fmt.Sprintf(`
			You are the MIA-Agentic-Data-Nervous-System (v3).
			Context: Analyze the following Knowledge Flow Graph (KFG) state for
			compliance anomalies (NIST 800-53) and data-drift.

			KFG STATE: %s

			OUTPUT: Provide a concise JSON action plan.
		`, kfgState)

		return a.model.GenerateContent(ctx, genai.Text(prompt))
	})

	// Handle Circuit Breaker Open state or other errors
	if err != nil {
		if errors.Is(err, resilience.ErrCircuitOpen) {
			span.AddEvent("Circuit Breaker Open - Returning Fail-Safe State")
			return `{"action": "STALL", "reason": "AI_SUBSYSTEM_LATENCY_CIRCUIT_OPEN"}`, nil
		}
		return "", fmt.Errorf("ai reasoning failed: %w", err)
	}

	// Cast and return result
	resp := result.(*genai.GenerateContentResponse)
	decision := fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0])

	span.SetAttributes(attribute.String("ai.decision", "COMPLETED"))
	return decision, nil
}