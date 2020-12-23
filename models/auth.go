package models

type Auth struct {
	Model
	Username string `gorm:"type:varchar(20)" json:"username"`
	Password string `gorm:"type:text" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

func CheckAuth(username, password string) bool {
	var auth Auth
	db.Model(&Auth{}).Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	return auth.ID > 0
}
