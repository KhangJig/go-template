package model

type Demo struct {
	ID   int64  `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
}
