package utils

import (
	"encoding/json"
	"fmt"
	"go-fiber-app/core/dto"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// FilterHintTransaction filters for transaction data based on area, province, hospital, and date.
func FilterHintTransaction(body map[string]string) (bson.M, error) {
	filter := bson.M{}

	if area, ok := body["area"]; ok && area != "" {
		num, err := strconv.Atoi(area)
		if err == nil {
			filter["area"] = int32(num) // ✅ แปลงเป็น int32 เพื่อให้ตรงกับ MongoDB
		}
	}

	if province, ok := body["province"]; ok && province != "" {
		filter["province"] = province
	}

	if hospital, ok := body["hospital"]; ok && hospital != "" {
		filter["hcode"] = hospital
	}

	if startDate, ok1 := body["start_date"]; ok1 {
		if endDate, ok2 := body["end_date"]; ok2 {
			layout := "2006-01-02T15:04:05.000-07:00"
			start, err1 := time.Parse(layout, fmt.Sprintf("%sT00:00:00.000+07:00", startDate))
			end, err2 := time.Parse(layout, fmt.Sprintf("%sT23:59:59.000+07:00", endDate))
			if err1 == nil && err2 == nil {
				filter["create_at"] = bson.M{
					"$gte": start,
					"$lte": end,
				}
			}
		}
	}

	return filter, nil
}

// FilterHintRegister filters for registration data including hospital, province, active status, and date range.
func FilterHintRegister(body map[string]string) (bson.M, error) {
	filter := bson.M{}

	if area, ok := body["area"]; ok && area != "" {
		num, err := strconv.Atoi(area)
		if err == nil {
			filter["area_main"] = int32(num)
		}
	}

	if province, ok := body["province"]; ok && province != "" {
		filter["province_main"] = province
	}

	if hospital, ok := body["hospital"]; ok && hospital != "" {
		filter["code_hospital_main"] = hospital
	}

	if startDate, ok1 := body["start_date"]; ok1 {
		if endDate, ok2 := body["end_date"]; ok2 {
			layout := "2006-01-02T15:04:05.000-07:00"
			start, err1 := time.Parse(layout, fmt.Sprintf("%sT00:00:00.000+07:00", startDate))
			end, err2 := time.Parse(layout, fmt.Sprintf("%sT23:59:59.000+07:00", endDate))
			if err1 == nil && err2 == nil {
				filter["register_date"] = bson.M{
					"$gte": start,
					"$lte": end,
				}
			}
		}
	}

	if active, ok := body["active"]; ok {
		if active == "inactive" {
			filter["change_right_date"] = bson.M{"$exists": true, "$type": "date"}
			filter["change_right_memo"] = bson.M{"$exists": true, "$ne": ""}
		} else if active == "active" {
			filter["$or"] = []bson.M{
				{"change_right_date": bson.M{"$exists": false}},
				{"change_right_date": bson.M{"$in": []interface{}{nil, ""}}},
				{"change_right_memo": bson.M{"$exists": false}},
				{"change_right_memo": bson.M{"$in": []interface{}{nil, ""}}},
			}
		}
	}

	return filter, nil
}

// ConvertFilterToMap converts HintfilterReq struct to a map[string]string
func ConvertFilterToMap(filter dto.HintfilterReq) map[string]string {
	result := map[string]string{}

	if filter.Area != nil {
		result["area"] = *filter.Area
	}
	if filter.Province != nil {
		result["province"] = *filter.Province
	}
	if filter.Hospital != nil {
		result["hcode"] = *filter.Hospital
	}
	
	

	c, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(c))
	return result
}
