package user

import (
	"codeReleaseSystem/app/models"
)

// User 用户模型
type User struct {
	models.BaseModel
	models.CommonTimestampsField
	Name     string `gorm:"column:name;not null;" json:"name"`
	Email    string `gorm:"column:email;not null;unique;" json:"email"`
	Phone    string `gorm:"column:phone;not null;unique;" json:"phone"`
	Password string `gorm:"column:password;not null;" json:"-"`
}
