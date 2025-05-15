package entities

import (
	"go-fiber-app/core/dto"
	"go-fiber-app/core/models"

	"github.com/gofiber/fiber/v2"
	// "net/http"
)

type HintRepository interface {
	QueryHintTransaction(filter dto.HintfilterReq) (results []models.Transaction, err error)
	QueryHintRegistration(filter dto.HintfilterReq) (results []models.Register, err error)
	CountHintRegister(filter dto.HintfilterReq) (int64, error)
	FindHintRegistrationBatch(filter dto.HintfilterReq, batchIndex int, batchSize int64) ([]models.Register, error)
	FindAllHintRegistration(filter dto.HintfilterReq) ([]models.Register, error)
	// QueryHintBudget(filter dto.HintfilterReq) (results []models.Budget, err error)
}

type HintService interface {
	GetHintTransaction(filter dto.HintfilterReq) (results []models.Transaction, err error)
	GetHintRegistration(filter dto.HintfilterReq) (results []models.Register, err error)
	GenerateRegistrationExcelReport(c *fiber.Ctx, body map[string]interface{}, filter dto.HintfilterReq) error
	// CountHintRegister(filter dto.HintfilterReq) (int64, error)
	// FindHintRegistrationBatch(filter dto.HintfilterReq, batchIndex int, batchSize int64) ([]models.Register, error)
}
