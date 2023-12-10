package main

import (
	"fmt"
	"ginchat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 设置数据库连接参数，包括 SSL 配置
	dsn := ""
	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("无法连接数据库：" + err.Error())
	}

	// Migrate the schema
	//db.AutoMigrate(&models.UserBasic{})

	// Read
	user := &models.UserBasic{}

	//Create
	//db.Create(user)

	fmt.Println(db.First(user, 1))
	//db.Model(user).Update("Password", "666")
}
