package models

type User struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"uniqueIndex;not null"`
	Password string `json:"password,omitempty" gorm:"not null"`
}
