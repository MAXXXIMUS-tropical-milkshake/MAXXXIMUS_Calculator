package calculator

type (
	Calculator struct {
		Expression string `query:"data" validate:"required"`
	}
)
