package model

import "log"

// 自动迁移
func migration() {
	err := DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&User{})

	if err != nil {
		log.Print("err")
	}
}
