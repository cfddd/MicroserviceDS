package model

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
	"user_service/config"
)

func TestInitDb(t *testing.T) {
	config.InitConfig()
	dsn := config.DbDsnInit()
	fmt.Println(dsn)
	fmt.Printf("host is : %v \n", viper.GetString("mysql.host"))
}
