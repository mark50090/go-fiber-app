package controller

import (
	// "fmt"

	"go-fiber-app/core/dto"
	"go-fiber-app/core/entities"
	"go-fiber-app/core/models"

	"github.com/gofiber/fiber/v2"
)

type HintController struct {
	service entities.HintService
}

func NewHandlerHint(service entities.HintService) *HintController {
	return &HintController{service: service}
}

func (ct *HintController) GetExample(c *fiber.Ctx) error {
	return c.JSON(models.Response{
		Message: "ok",
		Code:    fiber.StatusOK,
	})
}

func (ct *HintController) GetTransaction(c *fiber.Ctx) error {
	var filter dto.HintfilterReq
	if err := c.BodyParser(&filter); err != nil {
		return c.JSON(models.Response{
			Message: "ไม่สามารถแปลงข้อมูลที่รับเข้าได้",
			Code:    fiber.StatusBadRequest,
			Error:   err.Error(),
		})
	}

	results, err := ct.service.GetHintTransaction(filter)
	if err != nil {
		return c.JSON(models.Response{
			Message: "เกิดข้อผิดพลาดในการดึงข้อมูล",
			Code:    fiber.StatusInternalServerError,
			Error:   err.Error(),
		})
	}

	return c.JSON(models.Response{
		Message: "ดึงข้อมูลสำเร็จ",
		Code:    fiber.StatusOK,
		Data:    len(results),
	})
}

func (ct *HintController) GetHintRegistration(c *fiber.Ctx) error {
	var filter dto.HintfilterReq

	// ดึง query parameters ใส่ struct dto.HintfilterReq
	if err := c.QueryParser(&filter); err != nil {
		return c.JSON(models.Response{
			Message: "ไม่สามารถอ่าน query parameters ได้",
			Code:    fiber.StatusBadRequest,
			Error:   err.Error(),
		})
	}

	// ส่ง filter เข้า service
	results, err := ct.service.GetHintRegistration(filter)
	if err != nil {
		return c.JSON(models.Response{
			Message: "เกิดข้อผิดพลาดในการดึงข้อมูล",
			Code:    fiber.StatusInternalServerError,
			Error:   err.Error(),
		})
	}

	return c.JSON(models.Response{
		Message: "ดึงข้อมูลสำเร็จ",
		Code:    fiber.StatusOK,
		Data:    results,
	})
}

func (ct *HintController) ReportExcelRegistration(c *fiber.Ctx) error {
	var filter dto.HintfilterReq

	if err := c.BodyParser(&filter); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Message: "Invalid request body",
			Code:    fiber.StatusBadRequest,
			Error:   err.Error(),
		})
	}

	if err := ct.service.GenerateRegistrationExcelReport(c, filter); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Message: err.Error(),
			Code:    fiber.StatusBadRequest,
		})
	}

	return nil
}

func (ct *HintController) ReportExcelHiv(c *fiber.Ctx) error {
	var filter dto.HintfilterReq
	// b, _ := json.MarshalIndent(body, "", "  ")
	// fmt.Println(string(b))

	if err := c.BodyParser(&filter); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Message: "Invalid request body",
			Code:    fiber.StatusBadRequest,
			Error:   err.Error(),
		})
	}

	if err := ct.service.GenerateHivExcelReport(c, filter); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Message: err.Error(),
			Code:    fiber.StatusBadRequest,
		})
	}

	return nil
}
