package server

import (
	"fmt"
	"log"
	"net/http"
	"sign_go/config"
	"sign_go/handler"

	"github.com/gorilla/mux"
)

// 라우터 설정, 서버 시작
func StartServer() {

	// 설정 파일 로드
	_, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("설장 파일 로드 에러: %v", err)
	}

	// 라우터 설정
	router := mux.NewRouter()

	// POST 요청을 처리하는 엔드포인트 설정
	router.HandleFunc("/api/generate_key", handler.GenerateKeyHandler).Methods("POST")
	router.HandleFunc("/api/generate_sign", handler.GenerateSignHandler).Methods("POST")
	router.HandleFunc("/api/login", handler.MemberLoginHandler).Methods("POST") // MemberLogin 핸들러 추가

	// 서버 시작
	port := config.GConfig.Server.Port
	log.Printf("Starting server on : %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
