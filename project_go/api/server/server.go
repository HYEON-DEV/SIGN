package server

import (
	"SignProject/sign"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {
	log.Println("서버 실행")
	router := mux.NewRouter()

	// POST 요청을 처리하는 엔드포인트 설정
	router.HandleFunc("/api/generate_key", sign.GenerateKeyHandler).Methods("POST")
	router.HandleFunc("/api/generate_sign", sign.GenerateSignHandler).Methods("POST")

	log.Println("Starting server on :8081")
	http.ListenAndServe(":8081", router)
}
