package user

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	UserName string `gorm:"uniqueIndex"`
	Password string `gorm:"not null"`
}
