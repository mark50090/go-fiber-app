package models

type User struct {
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
	IsAdmin  bool   `json:"is_admin" bson:"is_admin"`
}
