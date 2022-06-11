package models

type Student struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     uint   `json:"age"`
}
