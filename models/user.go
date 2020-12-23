package models

type User struct {
	Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int `gorm:"type:int" json:"role"`
}
