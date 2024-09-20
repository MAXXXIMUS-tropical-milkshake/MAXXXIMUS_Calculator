package calculator

func ToResponse(ex float64) CalculatorResponse {
	return CalculatorResponse{
		EvaluatedValue: ex,
	}
}