package main

import (
	"SignProject/server"
	"SignProject/util"
)

func main() {
	// 로그 출력을 파일로 설정
	util.LogSetup()
	server.StartServer()
}
