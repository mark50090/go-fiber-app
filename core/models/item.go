package models

import "time"

type Item struct {
	Name  string  `bson:"name" json:"name"`
	Price float64 `bson:"price" json:"price"`
}

type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    any    `json:"data,omitempty"`
	Error   any    `json:"error,omitempty"`
}

type UnitCostSummaryFilter struct {
	StartDate time.Time `json:"start_date" query:"start_date"`
	EndDate   time.Time `json:"end_date" query:"end_date"`
	Hcode     string    `json:"hcode" query:"hcode"`
}

type UnitCostSummary struct {
	ICD    string  `json:"icd" bson:"icd"`
	Avg    float64 `json:"avg" bson:"avg"`
	NameEn string  `json:"name_en" bson:"name_en"`
	NameTh string  `json:"name_th" bson:"name_th"`
}
