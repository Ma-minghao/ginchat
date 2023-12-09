package main

import (
	"fmt"
	"ginchat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 设置数据库连接参数，包括 SSL 配置
	dsn := "root:180319@tcp(127.0.0.1:3306)/forlearning?charset=utf8mb4&parseTime=True&loc=Local"
	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("无法连接数据库：" + err.Error())
	}

	// Migrate the schema
	db.AutoMigrate(&models.UserBasic{})

	// Create
	user := &models.UserBasic{}

	db.Create(user)

	// Read
	fmt.Println(db.First(user,1))
	db.Model(user).Update("Password", "180319")
}
