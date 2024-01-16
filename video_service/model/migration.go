package model

import "log"

func migration() {
	// 自动迁移
	err := DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&Comment{})
	// Todo 判断error 写入日志
	if err != nil {
		log.Print("err")
	}
}
