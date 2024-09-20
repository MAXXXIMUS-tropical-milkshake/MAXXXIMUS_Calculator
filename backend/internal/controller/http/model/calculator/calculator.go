package calculator

type (
	Calculator struct {
		Expression string `query:"data" validate:"required"`
	}

	CalculatorResponse struct {
		EvaluatedValue float64 `json:"evaluated_value"`
	}
)
