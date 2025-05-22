package entities

import (
	"go-fiber-app/core/dto"
	"go-fiber-app/core/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	// "net/http"
)

type HintRepository interface {
	QueryHintTransaction(filter dto.HintfilterReq) (results []models.Transaction, err error)
	QueryHintRegistration(filter dto.HintfilterReq) (results []models.Register, err error)
	CountHintRegister(filter dto.HintfilterReq) (int64, error)
	FindHintRegistrationBatch(filter dto.HintfilterReq, batchIndex int, batchSize int64) ([]models.Register, error)
	FindAllHintRegistration(filter dto.HintfilterReq) ([]models.Register, error)
	FindDataHIVTemplate(filter dto.HintfilterReq) ([]models.Transaction, error)
	FindDataSTPTemplate(filter dto.HintfilterReq) (*mongo.Cursor, error)
	// QueryHintBudget(filter dto.HintfilterReq) (results []models.Budget, err error)
}

type HintService interface {
	GetHintTransaction(filter dto.HintfilterReq) (results []models.Transaction, err error)
	GetHintRegistration(filter dto.HintfilterReq) (results []models.Register, err error)
	GenerateRegistrationExcelReport(c *fiber.Ctx, filter dto.HintfilterReq) error
	GenerateHivExcelReport(c *fiber.Ctx, filter dto.HintfilterReq) error
	GenerateSTPExcelReport(c *fiber.Ctx, filter dto.HintfilterReq) error
	// CountHintRegister(filter dto.HintfilterReq) (int64, error)
	// FindHintRegistrationBatch(filter dto.HintfilterReq, batchIndex int, batchSize int64) ([]models.Register, error)
}
