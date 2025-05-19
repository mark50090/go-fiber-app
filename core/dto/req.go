package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QueryParser struct {
	// Area        *string `query:"area" bson:"area"`                 // HTTP query param: area, MongoDB field: area
	// Province    *string `query:"province" bson:"province"`         // HTTP query param: province, MongoDB field: province
	// Hcode       *string `query:"hcode" bson:"hcode"`               // HTTP query param: hcode, MongoDB field: hcode
	// Hname       *string `query:"hname" bson:"hname"`               // HTTP query param: hname, MongoDB field: hname
	// ServiceType *string `query:"service_type" bson:"service_type"` // HTTP query param: service_type, MongoDB field: service_type
	StartDate *string `query:"start_date" bson:"date"` // HTTP query param: start_date, MongoDB field: start_date
	EndDate   *string `query:"end_date" bson:"date"`   // HTTP query param: end_date, MongoDB field: end_date
	Inscl     *string `query:"inscl" bson:"inscl"`     // HTTP query param: inscl, MongoDB field: inscl
	// Status         *string `query:"status" bson:"status"`             // HTTP query param: status, MongoDB field: status
	// TransactionUID *string `query:"transaction_uid" bson:"transaction_uid"`
	// HN             *string `query:"hn" bson:"hn"`
	// AN             *string `query:"an" bson:"an"`
	// Code *string `query:"code" bson:"code"`
	// Search *string `query:"search"`
}

type QuerySearch struct {
	Search *string `query:"search"`
}

type QueryPage struct {
	Page  int `query:"page"`
	Limit int `query:"limit" bson:"limit"`
	Skip  int `query:"skip" bson:"skip"`
}

type CommentReq struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Message  string             `json:"message" bson:"message" validate:"required"`
	CreateAt time.Time          `json:"create_at" bson:"create_at"`
}

// type QueryItem struct {
// 	Name	string	`query:"name" bson:"name"`
// 	Price	int		`query:"price" bson:"price"`
// }

type HintfilterReq struct {
	Area     *string `query:"area" bson:"area"`
	Province *string `query:"province" bson:"province"`
	Hospital *string `query:"hcode" bson:"hcode"`
	// year int
	// Start_date *string `query:"start_date" bson:"date"`
	// End_date   *string `query:"end_date" bson:"date"`
}
