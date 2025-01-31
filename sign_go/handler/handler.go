package handler

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"test_hyeon/service"
	"test_hyeon/structs"
)

// ECDSA 개인키를 저장할 변수
var privateKey *ecdsa.PrivateKey

// 서비스 레이어를 초기화하는 변수
// var keyService *service.KeyService
var keyService service.KeyService

// 서비스 레이어 초기화, 서비스 구현체를 핸들러에 전달
func InitHandler(svc service.KeyService) {
	keyService = svc
}

/*
 *HTTP 요청을 처리하고, ECDSA 키를 생성하여 JSON 형식으로 변환한 후, DB에 저장
 */
func GenerateKeyHandler(w http.ResponseWriter, r *http.Request) {

	// ECDSA 키 생성
	var err error
	privateKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Printf("ECDSA 개인키 생성 실패: %v\n", err)
		return
	}
	log.Println("ECDSA 키 생성")

	// 키를 JSON 형식으로 변환

	jsonPrivateKey := &structs.PrivateKey{
		D:     privateKey.D,
		X:     privateKey.X,
		Y:     privateKey.Y,
		Curve: "P-256",
	}

	jsonPublicKey := &structs.PublicKey{
		X:     privateKey.PublicKey.X,
		Y:     privateKey.PublicKey.Y,
		Curve: "P-256",
	}

	// 키를 JSON으로 직렬화하고 바이트 배열로 변환 (json 구조체의 바이트 배열)

	privateKeyData, err := json.Marshal(jsonPrivateKey)
	if err != nil {
		log.Printf("개인키 JSON 직렬화 실패: %v", err)
		return
	}

	publicKeyData, err := json.Marshal(jsonPublicKey)
	if err != nil {
		log.Printf("공개키 JSON 직렬화 실패: %v", err)
		return
	}

	// 바이트 슬라이스([]byte)를 문자열로 변환하여 출력
	log.Printf("privatekey: %.30s...\n", string(privateKeyData))
	log.Printf("publickey: %.30s...\n", string(publicKeyData))

	// DB에 키 저장
	member := structs.Member{
		MemberID:   1, // ✅ 추후 실제 로그인 정보로 변경
		PrivateKey: sql.NullString{String: string(privateKeyData), Valid: true},
		PublicKey:  sql.NullString{String: string(publicKeyData), Valid: true},
	}

	err = keyService.SaveKeys(member)
	if err != nil {
		log.Printf("DB에 키 저장 실패: %v\n", err)
		return
	}

	// JSON 데이터를 파일에 저장

	// 사용자의 다운로드 폴더 경로 가져오기
	var downloadPath string
	if runtime.GOOS == "windows" {
		downloadPath = filepath.Join(os.Getenv("USERPROFILE"), "Downloads")
	} else {
		downloadPath = filepath.Join(os.Getenv("HOME"), "Downloads")
	}
	privateKeyFile := filepath.Join(downloadPath, "PrivateKey.json")
	publicKeyFile := filepath.Join(downloadPath, "PublicKey.json")

	err = os.WriteFile(privateKeyFile, privateKeyData, 0644)
	if err != nil {
		log.Printf("개인키 파일 저장 실패: %v", err)
		return
	}
	log.Println("개인키 파일 생성 - 성공")

	err = os.WriteFile(publicKeyFile, publicKeyData, 0644)
	if err != nil {
		log.Printf("공개키 파일 저장 실패: %v", err)
		return
	}
	log.Println("공개키 파일 생성 - 성공")

	// 응답으로 성공 메시지 반환
	response := map[string]string{
		"message":        "키 생성이 완료되었습니다.",
		"privateKeyFile": privateKeyFile,
		"publicKeyFile":  publicKeyFile,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

/*
 * 서명 생성 요청을 처리하는 핸들러 함수
 */
func GenerateSignHandler(w http.ResponseWriter, r *http.Request) {

	// 업로드된 파일 읽기
	// HTTP 요청에서 privateKeyFile이라는 필드 이름으로 전송된 파일을 읽어오는 함수
	// 여기서 "privateKeyFile"은 HTML <input type="file" name="privateKeyFile">에서 설정한 name 속성과 일치해야 한다.
	file, _, err := r.FormFile("input_privatekey")
	if err != nil {
		log.Printf("파일 읽기 실패 : %v\n", err)
		return
	}
	defer file.Close()

	// 파일 내용 바이트배열로 읽기
	privateKeyData, err := io.ReadAll(file)
	if err != nil {
		log.Printf("파일 내용 읽기 실패: %v\n", err)
		return
	}

	// 파일 내용 출력
	log.Printf("파일 내용 읽기 - 개인키: %.30s...\n", string(privateKeyData))

	/*
	 * 구

	 * 현

	 * 필

	 * 요
	 */

	// 응답으로 성공 메시지 반환
	response := map[string]string{
		"message": "서명 생성이 완료되었습니다.",
		// "signatureFile": signatureFile,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
