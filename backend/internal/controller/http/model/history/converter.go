package history

import "github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/internal/core"

func (p *GetAllHistoryParams) ToCoreGetAllHistoryParams() core.GetAllHistoryParams {
	return core.GetAllHistoryParams{
		Limit:  p.Limit,
		Offset: p.Offset,
	}
}
