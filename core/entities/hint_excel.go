package entities

import (
	"go-fiber-app/core/dto"
	"go-fiber-app/core/models"
)

type HintRepository interface {
	QueryHintTransaction(filter dto.HintfilterReq) (results []models.Transaction, err error)
	QueryHintRegistration(filter dto.HintfilterReq) (results []models.Register, err error)
	// QueryHintBudget(filter dto.HintfilterReq) (results []models.Budget, err error)
}

type HintService interface {
	GetHintTransaction(filter dto.HintfilterReq) (results []models.Transaction, err error)
	GetHintRegistration(filter dto.HintfilterReq) (results []models.Register, err error)
}
