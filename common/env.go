package common

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	accessKey string
	secretKey string
	appID     string
	appSecret string
	qqMapKey  string
	dsn       string
)

// GetEnv 获取env
func GetEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	accessKey = os.Getenv("QINIU_AK")
	secretKey = os.Getenv("QINIU_SK")
	appID = os.Getenv("APPID")
	appSecret = os.Getenv("APPSECRET")
	qqMapKey = os.Getenv("QQMAP_KEY")

	userName := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	database := os.Getenv("DATABASE")
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", userName, password, host, port, database)
	return
}

// GetDsn 获取 dsn
func GetDsn() string {
	return dsn
}

func GetQQMAPKey() string {
	return qqMapKey
}

// GetWxID 获取公众号key
func GetWxID() (key map[string]string) {
	key = make(map[string]string)
	key["appID"] = appID
	key["appSecret"] = appSecret
	return
}

// GetQiniuKey 获取七牛key
func GetQiniuKey() (key map[string]string) {
	key = make(map[string]string)
	key["accessKey"] = accessKey
	key["secretKey"] = secretKey
	return
}
