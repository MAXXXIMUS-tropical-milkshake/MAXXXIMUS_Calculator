package entity

import (
	"github.com/MAXXXIMUS-tropical-milkshake/MAXXXIMUS_Calculator/internal/core"
	"strings"
)

func (p *GetAllParams) FromUserRequest() core.GetAllParams {
	sortBy, sortOrder := "", ""
	if p.Sort != "" {
		splitSort := strings.Split(p.Sort, ",")

		sortBy = splitSort[0]
		sortOrder = splitSort[1]
	}

	return core.GetAllParams{
		SearchTerm: &p.SearchTerm,
		SortBy:     &sortBy,
		SortOrder:  &sortOrder,
	}
}
