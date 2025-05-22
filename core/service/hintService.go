package services

import (
	"encoding/json"
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
	// println(len(results))
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

func (s *HintService) GenerateRegistrationExcelReport(c *fiber.Ctx, filter dto.HintfilterReq) error {
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

func (s *HintService) GenerateHivExcelReport(c *fiber.Ctx, filter dto.HintfilterReq) error {
	// ดึงข้อมูลจาก repo
	docs, err := s.hintRepo.FindDataHIVTemplate(filter)
	// b, _ := json.MarshalIndent(docs, "", "  ")
	// fmt.Println(string(b))
	if err != nil {
		return err
	}

	f := excelize.NewFile()
	sheet := "Sheet1"
	f.NewSheet(sheet)
	utils.SetupHivSheet(f, sheet)

	rowNumber := 4
	index := 1

	for _, tx := range docs {
		rows := utils.SplitTransaction(tx) // ต้องมีฟังก์ชันนี้ใน utils
		for _, r := range rows {
			dobTh := utils.ToThaiDate(r.PatientDOB)
			createTh := utils.ToThaiDate(r.CreatedAt)
			age := utils.CalculateAge(r.PatientDOB)
			sexStr := utils.MapSex(r.Sex)

			f.SetCellValue(sheet, fmt.Sprintf("A%d", rowNumber), index)
			f.SetCellValue(sheet, fmt.Sprintf("B%d", rowNumber), r.HCode)
			f.SetCellValue(sheet, fmt.Sprintf("C%d", rowNumber), r.Hospital)
			f.SetCellValue(sheet, fmt.Sprintf("D%d", rowNumber), r.ServiceType)
			f.SetCellValue(sheet, fmt.Sprintf("E%d", rowNumber), r.HN)
			f.SetCellValue(sheet, fmt.Sprintf("F%d", rowNumber), r.AN)
			f.SetCellValue(sheet, fmt.Sprintf("G%d", rowNumber), r.PatientPID)
			f.SetCellValue(sheet, fmt.Sprintf("H%d", rowNumber), r.Fullname)
			f.SetCellValue(sheet, fmt.Sprintf("I%d", rowNumber), sexStr)
			f.SetCellValue(sheet, fmt.Sprintf("J%d", rowNumber), dobTh)
			f.SetCellValue(sheet, fmt.Sprintf("K%d", rowNumber), age)
			f.SetCellValue(sheet, fmt.Sprintf("L%d", rowNumber), createTh)
			f.SetCellValue(sheet, fmt.Sprintf("M%d", rowNumber), r.Code)
			f.SetCellValue(sheet, fmt.Sprintf("N%d", rowNumber), r.Name)
			f.SetCellValue(sheet, fmt.Sprintf("O%d", rowNumber), r.Amount)
			f.SetCellValue(sheet, fmt.Sprintf("P%d", rowNumber), r.Cost)
			f.SetCellValue(sheet, fmt.Sprintf("Q%d", rowNumber), r.SubmitAmount)

			rowNumber++
			index++
		}
	}

	utils.SetHivStyle(f, sheet)

	// แปลงเป็น buffer แล้วส่งเป็น Excel file
	buffer, err := f.WriteToBuffer()
	if err != nil {
		return err
	}

	c.Set(fiber.HeaderContentType, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Set(fiber.HeaderContentDisposition, `attachment; filename="Hint HIV.xlsx"`)
	c.Set(fiber.HeaderContentLength, strconv.Itoa(buffer.Len()))

	return c.SendStream(buffer)
}

func (s *HintService) GenerateSTPExcelReport(c *fiber.Ctx, filter dto.HintfilterReq) error {
	// ดึงข้อมูลจาก repo
	docs, err := s.hintRepo.FindDataSTPTemplate(filter)

	b, _ := json.MarshalIndent(docs, "", "  ")
	fmt.Println(string(b))
	if err != nil {
		return err
	}

	f := excelize.NewFile()
	sheet := "Sheet1"
	f.NewSheet(sheet)
	utils.SetupHivSheet(f, sheet)

	// rowNumber := 4
	// index := 1

	// for _, tx := range docs {
	// 	rows := utils.SplitTransaction(tx) // ต้องมีฟังก์ชันนี้ใน utils
	// 	for _, r := range rows {
	// 		dobTh := utils.ToThaiDate(r.PatientDOB)
	// 		createTh := utils.ToThaiDate(r.CreatedAt)
	// 		age := utils.CalculateAge(r.PatientDOB)
	// 		sexStr := utils.MapSex(r.Sex)

	// 		f.SetCellValue(sheet, fmt.Sprintf("A%d", rowNumber), index)
	// 		f.SetCellValue(sheet, fmt.Sprintf("B%d", rowNumber), r.HCode)
	// 		f.SetCellValue(sheet, fmt.Sprintf("C%d", rowNumber), r.Hospital)
	// 		f.SetCellValue(sheet, fmt.Sprintf("D%d", rowNumber), r.ServiceType)
	// 		f.SetCellValue(sheet, fmt.Sprintf("E%d", rowNumber), r.HN)
	// 		f.SetCellValue(sheet, fmt.Sprintf("F%d", rowNumber), r.AN)
	// 		f.SetCellValue(sheet, fmt.Sprintf("G%d", rowNumber), r.PatientPID)
	// 		f.SetCellValue(sheet, fmt.Sprintf("H%d", rowNumber), r.Fullname)
	// 		f.SetCellValue(sheet, fmt.Sprintf("I%d", rowNumber), sexStr)
	// 		f.SetCellValue(sheet, fmt.Sprintf("J%d", rowNumber), dobTh)
	// 		f.SetCellValue(sheet, fmt.Sprintf("K%d", rowNumber), age)
	// 		f.SetCellValue(sheet, fmt.Sprintf("L%d", rowNumber), createTh)
	// 		f.SetCellValue(sheet, fmt.Sprintf("M%d", rowNumber), r.Code)
	// 		f.SetCellValue(sheet, fmt.Sprintf("N%d", rowNumber), r.Name)
	// 		f.SetCellValue(sheet, fmt.Sprintf("O%d", rowNumber), r.Amount)
	// 		f.SetCellValue(sheet, fmt.Sprintf("P%d", rowNumber), r.Cost)
	// 		f.SetCellValue(sheet, fmt.Sprintf("Q%d", rowNumber), r.SubmitAmount)

	// 		rowNumber++
	// 		index++
	// 	}
	// }

	utils.SetHivStyle(f, sheet)

	// แปลงเป็น buffer แล้วส่งเป็น Excel file
	buffer, err := f.WriteToBuffer()
	if err != nil {
		return err
	}

	c.Set(fiber.HeaderContentType, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Set(fiber.HeaderContentDisposition, `attachment; filename="Hint HIV.xlsx"`)
	c.Set(fiber.HeaderContentLength, strconv.Itoa(buffer.Len()))

	return c.SendStream(buffer)
}
