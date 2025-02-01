package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"sign_go/service"
	"sign_go/util"
)

// 서비스 레이어를 초기화하는 변수
var MemberService service.MemberService

// 서비스 레이어 초기화, 서비스 구현체를 핸들러에 전달
func InitMemberHandler(svc service.MemberService) {
	MemberService = svc
}

// MemberLogin 핸들러 추가
func MemberLoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("MemberLoginHandler 시작")

	var loginRequest struct {
		UserID string `json:"user_id"`
		UserPW string `json:"user_pw"`
	}

	// 요청 바디에서 JSON 디코딩
	log.Println("요청 바디 디코딩 시작")
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		log.Println("요청 바디 디코딩 실패:", err)
		util.BadRequest(w, "잘못된 요청 형식")
		return
	}
	log.Println("요청 바디 디코딩 성공")

	// 서비스 레이어의 MemberLogin 메서드 호출
	log.Println("서비스 레이어의 MemberLogin 메서드 호출 시작")
	member, err := MemberService.MemberLogin(loginRequest.UserID, loginRequest.UserPW)
	if err != nil {
		log.Println("로그인 실패:", err)
		util.SendError(w, http.StatusUnauthorized, "로그인 실패: "+err.Error())
		return
	}
	log.Println("로그인 성공")

	// 응답으로 멤버 정보 반환
	log.Println("응답으로 멤버 정보 반환 시작")
	util.SendJSONOk(w, map[string]interface{}{"member": member})
	log.Println("응답으로 멤버 정보 반환 성공")
}
