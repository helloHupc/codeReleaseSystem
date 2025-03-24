package user

import (
	"codeReleaseSystem/pkg/database"
)

// IsEmailExist 判断邮箱是否已存在
func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(&User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// IsPhoneExist 判断手机号是否已存在
func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(&User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}