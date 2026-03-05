package logic

import (
	"context"
	"fmt"
	"cloud.google.com/go/vertexai/genai"
	"github.com/siralfbaez/mia-kfg-v3/pkg/observability"
	"go.opentelemetry.io/otel"
)

var (
	tracer = otel.Tracer("ai-agent")
)

type AgentLogic struct {
	client *genai.Client
	model  *genai.GenerativeModel
}

func NewAgentLogic(ctx context.Context, projectID, region string) (*AgentLogic, error) {
	client, err := genai.NewClient(ctx, projectID, region)
	if err != nil {
		return nil, err
	}
	// Using Gemini 1.5 Pro for complex reasoning over KFG state
	model := client.GenerativeModel("gemini-1.5-pro")
	return &AgentLogic{client: client, model: model}, nil
}

func (a *AgentLogic) ProcessEnrichedSignal(ctx context.Context, kfgState string) (string, error) {
	ctx, span := tracer.Start(ctx, "AI-Reasoning-Task")
	defer span.End()

	// System Prompt defining the "Nervous System" persona
	prompt := fmt.Sprintf(`
		You are the MIA-Agentic-Data-Nervous-System. 
		Current Knowledge Flow Graph State: %s
		Task: Analyze the stream and identify anomalies or compliance risks based on NIST 800-53.
	`, kfgState)

	resp, err := a.model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}

	// Staff Tip: Here you would parse the response into structured actions
	return fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0]), nil
}
