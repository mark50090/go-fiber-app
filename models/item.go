package models

type Item struct {
	Name  string  `bson:"name" json:"name"`
	Price float64 `bson:"price" json:"price"`
}
