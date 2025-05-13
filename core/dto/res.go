package dto

import "time"

type Pagination struct {
	CurrentPage  int64  `json:"current_page"`
	TotalPage    int64  `json:"total_page"`
	TotalItem    int64  `json:"total_items"`
	PageSize     int64  `json:"page_size"`
	NextPage     *int64 `json:"next_page"`
	PreviousPage *int64 `json:"previous_page"`
}

type UnitCostSummaryResponse struct {
	// Disease  string  `json:"disease" `
	ICD      string  `json:"icd"`
	NameEn   string  `json:"name_en"`
	NameTh   string  `json:"name_th"`
	TotalAvg float64 `json:"total"`
	Cost     float64 `json:"cost"`
	Laber    float64 `json:"laber"`
	Material float64 `json:"material" `
	Capital  float64 `json:"capital"`
	Direct   float64 `json:"direct"`
	Indirect float64 `json:"indirect" `
	Future   float64 `json:"future"`
}

type UnitCostSummaryTable struct {
	Disease    []UnitCostSummaryResponse `json:"disease"`
	Pagination Pagination                `json:"pagination"`
}

type InsclDetailResponse struct {
	Code   string `bson:"_id" json:"code"`
	Detail string `bson:"detail" json:"detail"`
}
type CommentRes struct {
	ID       string    `json:"id" bson:"_id"`
	Message  string    `json:"message" bson:"message"`
	CreateAt time.Time `json:"create_at" bson:"create_at"`
}

type ItemRespone struct {
	Name  string `query:"name" bson:"name"`
	Price int    `query:"price" bson:"price"`
}

// type CommentModel struct {
// 	ID       string `json:"id" bson:"_id"`
// 	Message  string `json:"message" bson:"message"`
// 	CreateAt string `json:"create_at" bson:"create_at"`
// }
