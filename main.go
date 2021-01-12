package main

import (
	"wozaizhao.com/api/server"
)

func main() {
	r := server.SetupRouter()

	r.Run(":9000")
}
