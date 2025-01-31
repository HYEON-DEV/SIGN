package server

import (
	"log"
	"net/http"
	"sign_go/handler"

	"github.com/gorilla/mux"
)

// 라우터 설정, 서버 시작
func StartServer() {

	// 라우터 설정
	router := mux.NewRouter()

	// POST 요청을 처리하는 엔드포인트 설정
	router.HandleFunc("/api/generate_key", handler.GenerateKeyHandler).Methods("POST")
	router.HandleFunc("/api/generate_sign", handler.GenerateSignHandler).Methods("POST")

	// 서버 시작
	log.Println("Starting server on :8081")
	http.ListenAndServe(":8081", router)
}
