package util

import (
	"encoding/json"
	"net/http"
	"sign_go/structs"
	"time"
)

// JSON 응답 생성 함수
func SendJSON(w http.ResponseWriter, status int, message string, data interface{}, err error) {
	// JSONResponse 구조체 생성, 초기화화
	response := structs.JSONResponse{
		Timestamp: time.Now(),
		Status:    status,
		Message:   message,
		Data:      data,
	}

	if err != nil {
		// 에러 정보를 담은 맵 생성
		errorData := map[string]interface{}{
			"error": err.Error(),
		}
		// Data 필드를 에러 정보로 덮어쓴다
		response.Data = errorData
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	// response 구조체를 JSON 형식으로 인코딩하여 응답으로 작성
	json.NewEncoder(w).Encode(response)
}

// 가변 인자를 사용한 JSON 응답 생성 함수
func SendJSONResponse(w http.ResponseWriter, status int, message string, args ...interface{}) {
	var data interface{}
	var err error

	// 가변 인자를 순회하면서 타입 스위치를 사용하여 data와 err 설정
	for _, arg := range args {
		// 인자의 타입 확인
		switch v := arg.(type) {
		case map[string]interface{}:
			data = v
		case error:
			err = v
		}
	}

	SendJSON(w, status, message, data, err)
}

// 200 OK 응답을 위한 함수
func SendJSONOk(w http.ResponseWriter, data map[string]interface{}) {
	SendJSONResponse(w, http.StatusOK, "OK", data)
}

// 200 OK 응답을 위한 함수 (데이터 없이)
func SendJSONOkNoData(w http.ResponseWriter) {
	SendJSONResponse(w, http.StatusOK, "OK")
}

// 에러 응답을 위한 함수
func SendError(w http.ResponseWriter, status int, message string) {
	SendJSONResponse(w, status, message)
}

// 400 Bad Request 응답을 위한 함수
func BadRequest(w http.ResponseWriter, message string) {
	SendError(w, http.StatusBadRequest, message)
}

// 500 Internal Server Error 응답을 위한 함수
func ServerError(w http.ResponseWriter, message string) {
	SendError(w, http.StatusInternalServerError, message)
}
