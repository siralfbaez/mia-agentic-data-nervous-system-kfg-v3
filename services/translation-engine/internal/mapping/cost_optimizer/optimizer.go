import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

var (
	tracer = otel.Tracer("cost-optimizer")
	meter  = otel.Meter("cost-optimizer")
)

func (s *StreamOptimizer) SelectBestPlan(ctx context.Context, plans []Plan) (Plan, error) {
	ctx, span := tracer.Start(ctx, "SelectBestPlan")
	defer span.End()

	// Metric: Track number of plans evaluated
	planCounter, _ := meter.Int64Counter("optimizer.plans_evaluated_total")
	planCounter.Add(ctx, int64(len(plans)))

	// ... (rest of your selection logic) ...

	span.SetAttributes(
		attribute.String("plan.id", bestPlan.ID),
		attribute.Float64("plan.cost", bestPlan.TotalCost),
	)
	
	return bestPlan, nil
}
