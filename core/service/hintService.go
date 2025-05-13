package services

import (
	"go-fiber-app/configs"
	"go-fiber-app/core/dto"
	"go-fiber-app/core/entities"
	"go-fiber-app/core/models"
)

type HintService struct {
	cfg configs.Config
	// cache        pkg.Cache
	hintRepo entities.HintRepository
	// logger       pkg.AppLog
}

func NewHintService(config configs.Config, repo entities.HintRepository) entities.HintService {
	return &HintService{
		cfg: config,
		// cache:        cache,
		hintRepo: repo,
		// logger:       log,
	}
}

func (s *HintService) GetHintTransaction(filter dto.HintfilterReq) (results []models.Transaction, err error) {
	// เรียกใช้งาน repository
	results, err = s.hintRepo.QueryHintTransaction(filter)
	println(len(results))
	// fmt.Println(results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (s *HintService) GetHintRegistration(filter dto.HintfilterReq) (results []models.Register, err error) {
	// เรียกใช้งาน repository
	results, err = s.hintRepo.QueryHintRegistration(filter)
	if err != nil {
		return nil, err
	}
	return results, nil
}
