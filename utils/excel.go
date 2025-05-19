package utils

import (
	"fmt"
	"go-fiber-app/core/models"
	"strconv"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SetupRegisSheet(f *excelize.File, sheetName string) {
	title := fmt.Sprintf("รายงานการขึ้นทะเบียนประกันสุขภาพบุคคลที่มีปัญหาสถานะและสิทธิ ข้อมูล ณ วันที่ %s", ToThaiDate(time.Now()))

	// Merge A1 ถึง P2
	f.MergeCell(sheetName, "A1", "P2")
	f.SetCellValue(sheetName, "A1", title)

	// ตั้งค่า style สำหรับ title
	styleTitle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true, Size: 18, Family: "TH SarabunPSK"},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
	})
	f.SetCellStyle(sheetName, "A1", "P2", styleTitle)

	// Header row (row 3)
	headers := []string{
		"ลำดับ", "เลขประจำตัว 13 หลัก", "ว/ด/ป เกิด", "อายุ", "เพศ",
		"คำนำหน้าชื่อ", "ชื่อ", "นามสกุล", "วันที่ลงทะเบียน",
		"รหัสหน่วยบริการหลัก", "ชื่อหน่วยบริการหลัก",
		"รหัสหน่วยบริการรอง", "ชื่อหน่วยบริการรอง", "จังหวัด",
		"วันที่เปลี่ยนสิทธิ", "หมายเหตุการเปลี่ยนสิทธิ",
	}

	f.SetSheetRow(sheetName, "A3", &[]interface{}{headers})

	for colIdx, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(colIdx+1, 3)
		f.SetCellValue(sheetName, cell, header)
	}

	// Style Header Row (row 3)
	headerStyle, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{Bold: true, Size: 14, Family: "TH SarabunPSK"},
		Alignment: &excelize.Alignment{
			Horizontal: "center", Vertical: "center",
		},
		// Fill: &excelize.Fill{
		// 	Type:    "pattern",
		// 	Color:   []string{"#9BC2E6"},
		// 	Pattern: 1,
		// },
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
	})

	f.SetCellStyle(sheetName, "A3", "P3", headerStyle)
	f.SetColWidth(sheetName, "A", "P", 30)
}

func SetupHivSheet(f *excelize.File, sheetName string) {
	title := fmt.Sprintf("ข้อมูลการรักษาผู้ป่วย HIV สิทธิประกันสุขภาพบุคคลที่มีปัญหาสถานะและสิทธิ ปีงบประมาณ พ.ศ. 2568")

	// Merge A1 ถึง P2
	f.MergeCell(sheetName, "A1", "Q2")
	f.SetCellValue(sheetName, "A1", title)

	// ตั้งค่า style สำหรับ title
	styleTitle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true, Size: 18, Family: "TH SarabunPSK"},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
	})
	f.SetCellStyle(sheetName, "A1", "Q2", styleTitle)

	// Header row (row 3)
	headers := []string{
		"ลำดับ", "รหัสหน่วยบริการ", "ชื่อสถานพยาบาล", "ประเภทผู้ป่วย", "HN", "AN", "PID",
		"ชื่อ-นามสกุล", "เพศ", "วันเดือนปีเกิด", "อายุ", "วันที่รับบริการ", "รหัสรายการที่ขอเบิก",
		"ชื่อรายการที่ขอเบิก", "จำนวนที่ขอเบิก", "ค่ารักษาที่เรียกเก็บ", "ยอดจ่ายชดเชย",
	}

	f.SetSheetRow(sheetName, "A3", &[]interface{}{headers})

	for colIdx, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(colIdx+1, 3)
		f.SetCellValue(sheetName, cell, header)
	}

	// Style Header Row (row 3)
	headerStyle, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{Bold: true, Size: 14, Family: "TH SarabunPSK"},
		Alignment: &excelize.Alignment{
			Horizontal: "center", Vertical: "center",
		},
		// Fill: &excelize.Fill{
		// 	Type:    "pattern",
		// 	Color:   []string{"#9BC2E6"},
		// 	Pattern: 1,
		// },
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
	})

	f.SetCellStyle(sheetName, "A3", "Q3", headerStyle)
	f.SetColWidth(sheetName, "A", "A", 10)
	f.SetColWidth(sheetName, "B", "Q", 20)
}

func SetRegisStyle(f *excelize.File, sheet string) {
	styleHeader, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold:   true,
			Size:   14,
			Family: "TH SarabunPSK",
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#9BC2E6"},
			Pattern: 1,
		},
		Border: []excelize.Border{
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "left", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
	})

	styleCell, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Size:   14,
			Family: "TH SarabunPSK",
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Border: []excelize.Border{
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "left", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
	})

	rows, _ := f.GetRows(sheet)
	for i := range rows {
		rowIdx := i + 1
		for col := 1; col <= len(rows[i]); col++ {
			cell, _ := excelize.CoordinatesToCellName(col, rowIdx)
			if rowIdx == 3 {
				f.SetCellStyle(sheet, cell, cell, styleHeader)
			} else {
				f.SetCellStyle(sheet, cell, cell, styleCell)
			}
		}
	}
}

func SetHivStyle(f *excelize.File, sheet string) {
	styleHeader, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold:   true,
			Size:   14,
			Family: "TH SarabunPSK",
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#9BC2E6"},
			Pattern: 1,
		},
		Border: []excelize.Border{
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "left", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
	})

	styleCell, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Size:   14,
			Family: "TH SarabunPSK",
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Border: []excelize.Border{
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "left", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
	})

	rows, _ := f.GetRows(sheet)
	for i := range rows {
		rowIdx := i + 1
		for col := 1; col <= len(rows[i]); col++ {
			cell, _ := excelize.CoordinatesToCellName(col, rowIdx)
			if rowIdx == 3 {
				f.SetCellStyle(sheet, cell, cell, styleHeader)
			} else {
				f.SetCellStyle(sheet, cell, cell, styleCell)
			}
		}
	}
}

func FormatRegisterRow(rowNumber int, data models.Register) []interface{} {
	dob := getDate(data.Dob)
	register := getDate(data.RegisterDate)
	changeDate := getDate(data.ChangeRightDate)

	memo := "-"
	if data.ChangeRightMemo != "" {
		memo = data.ChangeRightMemo
	}

	age := CalculateAge(dob)

	return []interface{}{
		rowNumber,
		data.Pid,
		ToThaiDate(dob),
		age,
		data.Sex,
		data.Title,
		data.Fname,
		data.Lname,
		ToThaiDate(register),
		data.CodeHospitalMain,
		data.HospitalMain,
		data.CodeHospitalSub,
		data.HospitalSub,
		data.ProvinceMain,
		ToThaiDate(changeDate),
		memo,
	}
}

func FormatHiv(rowNumber int, data models.Transaction) []interface{} {
	// fmt.Println(data.PatientDOB)
	// fmt.Println(getDate(data.PatientDOB))
	patient_dob := getDate(data.PatientDOB)

	fmt.Println(data.ADP)
	ADMTime := getDate(data.IPD.AdmDateTime)
	DSCTime := getDate(data.IPD.DscDateTime)

	// memo := "-"
	// if data.ChangeRightMemo != "" {
	// 	memo = data.ChangeRightMemo
	// }

	age := CalculateAge(patient_dob)

	return []interface{}{
		rowNumber,
		data.TransactionUID,
		ToThaiDate(patient_dob),
		age,
		data.Total,
		data.PatientPID,
		data.Fullname,
		data.HCode,
		data.Hospital,
		data.ServiceType,
		data.HN,
		data.AN,
		ToThaiDate(ADMTime),
		ToThaiDate(DSCTime),
		data.SubmitAmount,
	}
}

func ToThaiDate(t time.Time) string {
	if t.IsZero() {
		return "-"
	}
	thYear := t.Year()
	if thYear < 2400 {
		thYear += 543
	}
	return fmt.Sprintf("%02d/%02d/%d", t.Day(), t.Month(), thYear)
}

func CalculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()
	if now.YearDay() < dob.YearDay() {
		age--
	}
	return age
}

func getDate(v interface{}) time.Time {
	if v == nil {
		return time.Time{}
	}
	switch t := v.(type) {
	case time.Time:
		return t
	case primitive.DateTime:
		return t.Time()
	default:
		return time.Time{}
	}
}

// แปลง sex
func MapSex(code string) string {
	switch code {
	case "M":
		return "ชาย"
	case "F":
		return "หญิง"
	default:
		return "ไม่ระบุเพศ"
	}
}

type FlattenRow struct {
	models.Transaction
	Code   string
	Name   string
	Amount float64
	Cost   float64
}

// แยก adp และ dru
func SplitTransaction(tx models.Transaction) []FlattenRow {
	var rows []FlattenRow

	for _, adp := range tx.ADP {
		qty, _ := strconv.ParseFloat(adp.Qty, 64)
		rate, _ := strconv.ParseFloat(adp.Rate, 64)
		cost := qty * rate

		name := strings.TrimSpace(adp.Name)
		if name == "" {
			name = "-"
		}

		rows = append(rows, FlattenRow{
			Transaction: tx,
			Code:        adp.Code,
			Name:        name,
			Amount:      qty,
			Cost:        cost,
		})
	}

	for _, dru := range tx.DRU {
		amount, _ := strconv.ParseFloat(dru.Amount, 64)
		price, _ := strconv.ParseFloat(dru.Price, 64)
		cost := amount * price

		name := strings.TrimSpace(dru.DIDName)
		if name == "" {
			name = "-"
		}

		rows = append(rows, FlattenRow{
			Transaction: tx,
			Code:        dru.DID,
			Name:        name,
			Amount:      amount,
			Cost:        cost,
		})
	}

	return rows
}
