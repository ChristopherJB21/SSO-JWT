package app

import (
	"sso-jwt/helper"
	"sso-jwt/model/user"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := viper.GetString("database.dsn")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)

	sqlDB, err := db.DB()
	helper.PanicIfError(err)

	sqlDB.SetMaxIdleConns(viper.GetInt("database.maxidleconns"))
	sqlDB.SetMaxOpenConns(viper.GetInt("database.maxopenconns"))
	sqlDB.SetConnMaxLifetime(viper.GetDuration("database.connmaxlifetime") * time.Minute)
	sqlDB.SetConnMaxIdleTime(viper.GetDuration("database.connmaxidletime") * time.Minute)

	db.AutoMigrate(&user.User{})

	return db
}
