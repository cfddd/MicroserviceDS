package config

import (
	"fmt"
	"github.com/spf13/viper"
	"path"
	"runtime"
	"testing"
)

func TestInitConfig(t *testing.T) {
	_, filepath, _, _ := runtime.Caller(0)
	currentDir := path.Dir(filepath)
	fmt.Println(currentDir)
}

func TestDbDsnInit(t *testing.T) {
	InitConfig()
	fmt.Printf("host is : %v \n", viper.GetString("mysql.host"))
	dsn := DbDsnInit()
	fmt.Printf("dns is : %v", dsn)
}
