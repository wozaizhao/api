package common

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	accessKey string
	secretKey string
	appID     string
	appSecret string
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

	return
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
