package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"sign_go/service"
	"sign_go/util"
	"time"
)

// 서비스 레이어를 초기화하는 변수
var MemberService service.MemberService

// 서비스 레이어 초기화, 서비스 구현체를 핸들러에 전달
func InitMemberHandler(svc service.MemberService) {
	MemberService = svc
}

// MemberLoginHandler - 로그인 핸들러
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

	// 세션 ID 생성 ( UUID 사용 )
	sessionID := util.GenerateSessionID()

	// 세션 데이터를 DB에 저장
	expiresAt := time.Now().Add(24 * time.Hour) // 세션 만료 시간: 24시간
	err = MemberService.SaveSession(sessionID, member.MemberID, member.Name, member.UserID, expiresAt)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "세션 저장 실패")
		return
	}

	// 세션 ID를 클라이언트에게 반환
	util.SendJSONOk(w, map[string]interface{}{"session_id": sessionID})

}

// MemberLogoutHandler - 로그아웃 핸들러
func MemberLogoutHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("MemberLogoutHandler 시작")

	var logoutRequest struct {
		SessionID string `json:"session_id"`
	}

	// 요청 바디에서 JSON 디코딩
	err := json.NewDecoder(r.Body).Decode(&logoutRequest)
	if err != nil {
		util.BadRequest(w, "잘못된 요청 형식")
		return
	}

	// 세션 데이터를 DB에서 삭제
	err = MemberService.DeleteSession(logoutRequest.SessionID)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "로그아웃 실패")
		return
	}

	util.SendJSONOk(w, map[string]interface{}{"message": "로그아웃 성공"})
}
