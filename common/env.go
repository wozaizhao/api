package common

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	accessKey   string
	secretKey   string
	appID       string
	appSecret   string
	mongodbHost string
	mongodbPort string
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
	mongodbHost = os.Getenv("MONGODB_HOST")
	mongodbPort = os.Getenv("MONGODB_PORT")

	return
}

// GetWxID 获取公众号key
func GetWxID() (key map[string]string) {
	key = make(map[string]string)
	key["appID"] = appID
	key["appSecret"] = appSecret
	return
}

// GetMongodbURL 获取mongodb连接地址
func GetMongodbURL() string {
	return "mongodb://" + mongodbHost + ":" + mongodbPort
}

// GetQiniuKey 获取七牛key
func GetQiniuKey() (key map[string]string) {
	key = make(map[string]string)
	key["accessKey"] = accessKey
	key["secretKey"] = secretKey
	return
}
