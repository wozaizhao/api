package main

import (
	"wozaizhao.com/api/common"
	"wozaizhao.com/api/controllers/user"
	"wozaizhao.com/api/server"
)

func main() {
	r := server.SetupRouter()
	common.GetEnv()
	user.Init()
	r.Run(":9000")
}
