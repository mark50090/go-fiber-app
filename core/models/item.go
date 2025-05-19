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

type HivExcel struct {
	TransactionUID string  `json:"transaction_uid" bson:"transaction_uid"`
	PatientDOB     any     `json:"patient_dob" bson:"patient_dob"`
	Total          string  `json:"total" bson:"total"`
	PatientPID     string  `json:"patient_pid" bson:"patient_pid"`
	Fullname       string  `bson:"fullname" json:"fullname"`
	HCode          string  `bson:"hcode" json:"hcode"`
	Hospital       string  `bson:"hospital" json:"hospital"`
	ServiceType    string  `bson:"service_type" json:"service_type"`
	HN             string  `bson:"hn" json:"hn"`
	AN             string  `bson:"an" json:"an"`
	ADMTime        any     `bson:"adm_date_time" json:"adm_date_time"`
	DSCTime        any     `bson:"dsc_date_time" json:"dsc_date_time"`
	SubmitAmount   float64 `bson:"submit_amount" json:"submit_amount"`
}
