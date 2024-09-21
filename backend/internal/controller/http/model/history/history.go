package history

type GetAllHistoryParams struct {
	Limit  int `query:"limit" validate:"gt=0"`
	Offset int `query:"limit" validate:"gte=0"`
}
