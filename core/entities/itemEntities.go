package entities

import (
	"go-fiber-app/core/dto"
	"go-fiber-app/core/models"
)

type UnitCostRepository interface {
	QueryUnitCostSummary(filter dto.QueryParser, QueryPage dto.QueryPage, querySearch dto.QuerySearch) (results []models.UnitCostSummary, total int64, err error)
	QueryInsclDetail() (results []dto.InsclDetailResponse, err error)
	QueryItemData() (results []dto.ItemRespone, err error)
}

type UnitCostServices interface {
	GetUnitCostSummary(filter dto.QueryParser, QueryPage dto.QueryPage, QuerySearch dto.QuerySearch) (results *dto.UnitCostSummaryTable, err error)
	GetInsclDetail() (results []dto.InsclDetailResponse, err error)
	GetItemData() (results []dto.ItemRespone, err error)
}
