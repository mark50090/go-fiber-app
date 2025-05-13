package services

import (
	"go-fiber-app/configs"
	"go-fiber-app/core/dto"
	"go-fiber-app/core/entities"
	"go-fiber-app/utils"
)

type UnitCostService struct {
	cfg configs.Config
	// cache        pkg.Cache
	unitCostRepo entities.UnitCostRepository
	// logger       pkg.AppLog
}

func NewUnitCostService(config configs.Config, repo entities.UnitCostRepository) entities.UnitCostServices {
	return &UnitCostService{
		cfg: config,
		// cache:        cache,
		unitCostRepo: repo,
		// logger:       log,
	}
}

func (s *UnitCostService) GetUnitCostSummary(filter dto.QueryParser, QueryPage dto.QueryPage, QuerySearch dto.QuerySearch) (results *dto.UnitCostSummaryTable, err error) {

	// totalItems := int64(0)
	resQuery, totalItems, err := s.unitCostRepo.QueryUnitCostSummary(filter, QueryPage, QuerySearch)
	if err != nil {
		return nil, err
	}
	results = &dto.UnitCostSummaryTable{}
	for _, res := range resQuery {
		cost := res.Avg * 0.75
		laber := cost * 0.4
		material := cost * 0.25
		capital := cost * 0.15
		direct := laber + material + capital
		indirect := cost * 0.2
		future := res.Avg - cost
		results.Disease = append(results.Disease, dto.UnitCostSummaryResponse{
			ICD:      res.ICD,
			NameEn:   res.NameEn,
			NameTh:   res.NameTh,
			TotalAvg: res.Avg,
			Cost:     cost,
			Laber:    laber,
			Material: material,
			Capital:  capital,
			Direct:   direct,
			Indirect: indirect,
			Future:   future,
		})
	}

	results.Pagination = utils.CalculatePagination(totalItems, QueryPage)
	// ดึงข้อมูลล่วงหน้าสำหรับหน้าถัดไป 2-3 หน้า
	for i := 1; i <= 3; i++ {
		nextPageNum := QueryPage.Page + i
		if nextPageNum <= int(results.Pagination.TotalPage) {
			nextPage := dto.QueryPage{
				Page:  nextPageNum,
				Limit: QueryPage.Limit,
			}

			// ดึงข้อมูลหน้าถัดไปแบบไม่บล็อกการทำงาน
			go func(page dto.QueryPage) {
				_, _, _ = s.unitCostRepo.QueryUnitCostSummary(filter, page, QuerySearch)
			}(nextPage)
		}
	}
	return results, nil
}

func (s *UnitCostService) GetInsclDetail() (results []dto.InsclDetailResponse, err error) {
	// เรียกใช้งาน repository
	results, err = s.unitCostRepo.QueryInsclDetail()
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (s *UnitCostService) GetItemData() (results []dto.ItemRespone, err error) {
	// เรียกใช้งาน repository
	results, err = s.unitCostRepo.QueryItemData()
	if err != nil {
		return nil, err
	}
	return results, nil
}
