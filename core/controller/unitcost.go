package controller

import (
	"go-fiber-app/core/dto"
	"go-fiber-app/core/entities"
	"go-fiber-app/core/models"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UnitCostController struct {
	service entities.UnitCostServices
}

func NewHandlerUnitCost(service entities.UnitCostServices) *UnitCostController {
	return &UnitCostController{service: service}
}

func (ct *UnitCostController) GetExample(c *fiber.Ctx) error {
	return c.JSON(models.Response{
		Message: "ok",
		Code:    fiber.StatusOK,
	})
}

func (ct *UnitCostController) GetUnitCostSummary(c *fiber.Ctx) error {
	// สร้างตัวแปรสำหรับเก็บพารามิเตอร์การค้นหา

	startDate := time.Date(2024, 4, 1, 0, 0, 0, 0, time.Local).Format("2006-01-02 00:00:00")
	endDate := time.Date(2024, 9, 30, 0, 0, 0, 0, time.Local).Format("2006-01-02 15:04:05")
	filter := dto.QueryParser{
		StartDate: &startDate,
		EndDate:   &endDate,
	}
	queryPage := dto.QueryPage{
		Limit: 10, // ค่าเริ่มต้น
		Page:  1,  // ค่าเริ่มต้น
	}
	querySearch := dto.QuerySearch{}

	// แยกการตรวจสอบข้อผิดพลาดสำหรับแต่ละการแปลงค่า
	if err := c.QueryParser(&filter); err != nil {
		return c.JSON(models.Response{
			Message: "ข้อผิดพลาดในการแปลงค่าพารามิเตอร์ filter",
			Code:    fiber.StatusBadRequest,
			Error:   err.Error(),
		})
	}

	if err := c.QueryParser(&queryPage); err != nil {
		return c.JSON(models.Response{
			Message: "ข้อผิดพลาดในการแปลงค่าพารามิเตอร์หน้า",
			Code:    fiber.StatusBadRequest,
			Error:   err.Error(),
		})
	}

	if err := c.QueryParser(&querySearch); err != nil {
		return c.JSON(models.Response{
			Message: "ข้อผิดพลาดในการแปลงค่าพารามิเตอร์การค้นหา",
			Code:    fiber.StatusBadRequest,
			Error:   err.Error(),
		})
	}
	if querySearch.Search != nil && *querySearch.Search != "" {
		replace := strings.ReplaceAll(*querySearch.Search, ".", "")
		querySearch.Search = &replace
	}

	// ดึงข้อมูลจาก service
	results, err := ct.service.GetUnitCostSummary(filter, queryPage, querySearch)
	if err != nil {
		return c.JSON(models.Response{
			Message: "เกิดข้อผิดพลาดในการดึงข้อมูล",
			Code:    fiber.StatusInternalServerError,
			Error:   err.Error(),
		})
	}

	// ส่งผลลัพธ์กลับไป
	return c.JSON(models.Response{
		Message: "ดึงข้อมูลสำเร็จ",
		Code:    fiber.StatusOK,
		Data:    results,
	})
}

func (ct *UnitCostController) GetInsclDetail(c *fiber.Ctx) error {

	results, err := ct.service.GetInsclDetail()
	if err != nil {
		return c.JSON(models.Response{
			Message: "เกิดข้อผิดพลาดในการดึงข้อมูล",
			Code:    fiber.StatusInternalServerError,
			Error:   err.Error(),
		})
	}
	// ส่งผลลัพธ์กลับไป
	return c.JSON(models.Response{
		Message: "ดึงข้อมูลสำเร็จ",
		Code:    fiber.StatusOK,
		Data:    results,
	})
}

func (ct *UnitCostController) GetItemData(c *fiber.Ctx) error {
	results, err := ct.service.GetItemData()
	if err != nil {
		return c.JSON(models.Response{
			Message: "เกิดข้อผิดพลาดในการดึงข้อมูล",
			Code:    fiber.StatusInternalServerError,
			Error:   err.Error(),
		})
	}
	// ส่งผลลัพธ์กลับไป
	return c.JSON(models.Response{
		Message: "ดึงข้อมูลสำเร็จ",
		Code:    fiber.StatusOK,
		Data:    results,
	})
}
