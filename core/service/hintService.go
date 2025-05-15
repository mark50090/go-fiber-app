package services

import (
	"fmt"
	"go-fiber-app/configs"
	"go-fiber-app/core/dto"
	"go-fiber-app/core/entities"
	"go-fiber-app/core/models"
	"go-fiber-app/utils"
	"strconv"

	// "net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
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

// func (s *HintService) GenerateRegistrationExcelReport(c *fiber.Ctx, body map[string]interface{}, filter dto.HintfilterReq) error {
// 	totalCount, err := s.hintRepo.CountHintRegister(filter)
// 	if err != nil {
// 		return fmt.Errorf("Failed to count documents: %v", err)
// 	}
// 	if totalCount >= 100000 {
// 		c.Status(fiber.StatusBadRequest)
// 		return c.JSON(fiber.Map{
// 			"total":   totalCount,
// 			"limit":   50000,
// 			"status":  false,
// 			"message": "ข้อมูลที่ต้องการดาวน์โหลดมีขนาดใหญ่เกินไป รอการปรับปรุงหรือแจ้งเจ้าหน้าที่",
// 		})
// 	}
// 	f := excelize.NewFile()
// 	sheet := "Sheet1"
// 	f.NewSheet(sheet)
// 	utils.SetupRegisSheet(f, sheet)
// 	batchSize := int64(10000)
// 	sumCount := int((totalCount + batchSize - 1) / batchSize)
// 	rowNumber := 1
// 	for i := 0; i < sumCount; i++ {
// 		docs, err := s.hintRepo.FindHintRegistrationBatch(filter, i, batchSize)
// 		if err != nil {
// 			return fmt.Errorf("Query error at batch %d: %v", i, err)
// 		}
// 		for _, data := range docs {
// 			rowNumber++
// 			row := utils.FormatRegisterRow(rowNumber, data)
// 			cell, _ := excelize.CoordinatesToCellName(1, rowNumber)
// 			f.SetSheetRow(sheet, cell, &row)
// 		}
// 	}
// 	utils.SetRegisStyle(f, sheet)
// 	buffer, err := f.WriteToBuffer()
// 	if err != nil {
// 		return err
// 	}
// 	c.Set(fiber.HeaderContentType, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
// 	c.Set(fiber.HeaderContentDisposition, `attachment; filename="Hint Dashboard.xlsx"`)
// 	c.Set(fiber.HeaderContentLength, strconv.Itoa(buffer.Len()))
// 	if err := c.SendStream(buffer); err != nil {
// 		return err
// 	}
// 	return nil
// }

func (s *HintService) GenerateRegistrationExcelReport(c *fiber.Ctx, body map[string]interface{}, filter dto.HintfilterReq) error {
	// 1. ดึงข้อมูลทั้งหมด (ไม่มี count, skip, หรือ limit)
	docs, err := s.hintRepo.FindAllHintRegistration(filter)
	if err != nil {
		return fmt.Errorf("Failed to fetch documents: %v", err)
	}

	// 2. เตรียมไฟล์ Excel
	f := excelize.NewFile()
	sheet := "Sheet1"
	f.NewSheet(sheet)
	utils.SetupRegisSheet(f, sheet)

	rowIndex := 4
	for i, data := range docs {
		row := utils.FormatRegisterRow(i+1, data)
		cell, _ := excelize.CoordinatesToCellName(1, rowIndex)
		f.SetSheetRow(sheet, cell, &row)
		rowIndex++
	}

	utils.SetRegisStyle(f, sheet)

	// 3. เขียนไฟล์ลง buffer แล้วส่งออก
	buffer, err := f.WriteToBuffer()
	if err != nil {
		return err
	}

	c.Set(fiber.HeaderContentType, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Set(fiber.HeaderContentDisposition, `attachment; filename="Hint Dashboard.xlsx"`)
	c.Set(fiber.HeaderContentLength, strconv.Itoa(buffer.Len()))

	return c.SendStream(buffer)
}
