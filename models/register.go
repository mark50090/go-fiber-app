package models

import "time"

type Register struct {
	Batch            string    `bson:"batch" json:"batch"`
	AreaMain         int       `bson:"area_main" json:"area_main"`
	AreaSub          int       `bson:"area_sub" json:"area_sub"`
	CodeHospitalMain string    `bson:"code_hospital_main" json:"code_hospital_main"`
	CodeHospitalSub  string    `bson:"code_hospital_sub" json:"code_hospital_sub"`
	HospitalMain     string    `bson:"hospital_main" json:"hospital_main"`
	HospitalSub      string    `bson:"hospital_sub" json:"hospital_sub"`
	ProvinceMain     string    `bson:"province_main" json:"province_main"`
	ProvinceSub      string    `bson:"province_sub" json:"province_sub"`
	Pid              string    `bson:"pid" json:"pid"`
	Dob              time.Time `bson:"dob" json:"dob"`
	Sex              string    `bson:"sex" json:"sex"`
	Title            string    `bson:"title" json:"title"`
	Fname            string    `bson:"fname" json:"fname"`
	Lname            string    `bson:"lname" json:"lname"`
	Fullname         string    `bson:"fullname" json:"fullname"`
	RegisterDate     time.Time `bson:"register_date" json:"register_date"`
	Status           string    `bson:"status" json:"status"`
	TypeHospitalMain string    `bson:"type_hospital_main" json:"type_hospital_main"`
	ChangeRightDate  time.Time `bson:"change_right_date" json:"change_right_date"`
	ChangeRightMemo  string    `bson:"change_right_memo" json:"change_right_memo"`
	CreatedAt        time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt        time.Time `bson:"updated_at" json:"updated_at"`
}
