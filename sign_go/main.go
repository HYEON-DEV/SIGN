package main

import (
	"log"
	"sign_go/config"
	"sign_go/db"
	"sign_go/handler"
	"sign_go/server"
	"sign_go/util"
)

func main() {
	util.LogSetup()
	util.Enterlog("main")
	defer util.Leavelog("main")

	// 설정 로드
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// DB 연결 초기화
	// 데이터베이스 연결 문자열(DSN)을 설정에서 가져오기기
	dsn := config.MySQLDSN()
	// 데이터베이스 연결 초기화, 객체 생성성
	mysqlDAO, err := db.NewMySQLDAO(dsn)
	if err != nil {
		log.Fatalf("[ERROR] Failed to connect MySQL: %v", err)
	}
	defer func() { // 프로그램 종료 시 DB를 닫도록 설정
		if mysqlDAO != nil {
			if err := mysqlDAO.Close(); err != nil {
				log.Printf("[WARN] Failed to close mySQL connection: %v", err)
			}
		}
	}()
	log.Println("DB 연결")

	// 핸들러 초기화
	// 서비스 레이어를 초기화, DB 접근 객체(mysqlDAO)를 핸들러에 전달
	handler.InitHandler(mysqlDAO)

	// 서버 시작
	server.StartServer()

}
