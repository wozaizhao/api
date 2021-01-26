package main

import (
	"wozaizhao.com/api/common"
	"wozaizhao.com/api/controllers/user"
	"wozaizhao.com/api/models"
	"wozaizhao.com/api/server"
)

func main() {
	r := server.SetupRouter()
	common.GetEnv()
	user.Init()
	models.OpenDB()
	r.Run(":9000")
}
