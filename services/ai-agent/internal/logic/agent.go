package logic

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/vertexai/genai"
	"github.com/siralfbaez/mia-agentic-data-nervous-system-kfg-v3/pkg/resilience"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

var (
	tracer = otel.Tracer("ai-agent")
)

type AgentLogic struct {
	client     *genai.Client
	model      *genai.GenerativeModel
	resilience *resilience.Policy
}

func NewAgentLogic(ctx context.Context, projectID, region string) (*AgentLogic, error) {
	client, err := genai.NewClient(ctx, projectID, region)
	if err != nil {
		return nil, fmt.Errorf("failed to init vertex ai: %w", err)
	}

	model := client.GenerativeModel("gemini-1.5-pro")
	policy := resilience.NewResiliencePolicy("gemini-reasoning-cb")

	return &AgentLogic{
		client:     client,
		model:      model,
		resilience: policy,
	}, nil
}

func (a *AgentLogic) ProcessEnrichedSignal(ctx context.Context, kfgState string) (string, error) {
	ctx, span := tracer.Start(ctx, "AI-Reasoning-Task")
	defer span.End()

	span.SetAttributes(attribute.Int("state.length", len(kfgState)))

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

	if err != nil {
		if errors.Is(err, resilience.ErrCircuitOpen) {
			span.AddEvent("Circuit Breaker Open - Returning Fail-Safe State")
			return `{"action": "STALL", "reason": "AI_SUBSYSTEM_LATENCY_CIRCUIT_OPEN"}`, nil
		}
		return "", fmt.Errorf("ai reasoning failed: %w", err)
	}

	resp := result.(*genai.GenerateContentResponse)
	decision := fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0])

	span.SetAttributes(attribute.String("ai.decision", "COMPLETED"))
	return decision, nil
} // <--- THIS BRACE WAS MISSING

// StartAgent is the entry point called by cmd/main.go
func StartAgent(ctx context.Context) error {
	log.Println("Initializing AI Agent Logic...")

	projectID := os.Getenv("GCP_PROJECT_ID")
	region := os.Getenv("GCP_REGION")
	if projectID == "" {
		projectID = "your-project-id"
	}
	if region == "" {
		region = "us-central1"
	}

	_, err := NewAgentLogic(ctx, projectID, region)
	if err != nil {
		return fmt.Errorf("failed to initialize agent logic: %w", err)
	}

	log.Println("AI Agent Logic initialized. Waiting for KFG signals...")

	<-ctx.Done()

	log.Println("Stopping AI Agent Logic...")
	return nil
}
