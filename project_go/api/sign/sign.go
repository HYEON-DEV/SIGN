package sign

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

// ECDSA 키를 JSON 형식으로 변환하기 위한 구조체 정의

// PrivateKey 구조체 정의
type PrivateKey struct {
	D     *big.Int `json:"d"`
	X     *big.Int `json:"x"`
	Y     *big.Int `json:"y"`
	Curve string   `json:"curve"`
}

// PublicKey 구조체 정의
type PublicKey struct {
	X     *big.Int `json:"x"`
	Y     *big.Int `json:"y"`
	Curve string   `json:"curve"`
}

var privateKey *ecdsa.PrivateKey // ECDSA 개인키를 저장할 변수

/*
 * 키 생성 요청을 처리하는 핸들러 함수
 */
func GenerateKeyHandler(w http.ResponseWriter, r *http.Request) {

	// ECDSA 키 생성
	var err error
	privateKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Printf("ECDSA 개인키 생성 실패: %v", err)
		http.Error(w, "키 생성 실패", http.StatusInternalServerError)
		return
	}
	log.Println("ECDSA 키 생성")

	// 키를 JSON 형식으로 변환

	jsonPrivateKey := &PrivateKey{
		D:     privateKey.D,
		X:     privateKey.X,
		Y:     privateKey.Y,
		Curve: "P-256",
	}

	jsonPublicKey := &PublicKey{
		X:     privateKey.PublicKey.X,
		Y:     privateKey.PublicKey.Y,
		Curve: "P-256",
	}

	// 키를 JSON으로 직렬화하고 바이트 배열로 변환 (json 구조체의 바이트 배열)

	privateKeyData, err := json.Marshal(jsonPrivateKey)
	if err != nil {
		log.Printf("개인키 JSON 직렬화 실패: %v", err)
		http.Error(w, "개인키 JSON 직렬화 실패", http.StatusInternalServerError)
		return
	}

	publicKeyData, err := json.Marshal(jsonPublicKey)
	if err != nil {
		log.Printf("공개키 JSON 직렬화 실패: %v", err)
		http.Error(w, "공개키 JSON 직렬화 실패", http.StatusInternalServerError)
		return
	}

	/*
	 *
	 * ✅	DB에 개인키, 공개키 저장		✅
	 *
	 */

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
		http.Error(w, "개인키 파일 저장 실패", http.StatusInternalServerError)
		return
	}
	log.Println("개인키 파일 생성 - 성공")

	err = os.WriteFile(publicKeyFile, publicKeyData, 0644)
	if err != nil {
		log.Printf("공개키 파일 저장 실패: %v", err)
		http.Error(w, "공개키 파일 저장 실패", http.StatusInternalServerError)
		return
	}
	log.Println("공개키 파일 생성 - 성공")

	// 응답으로 성공 메시지 반환
	response := map[string]string{ // ✅이 response 는 어디에에
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
		log.Printf("파일 읽기 실패 : %v", err)
		http.Error(w, "파일 읽기 실패", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// 파일 내용 바이트배열로 읽기
	privateKeyData, err := io.ReadAll(file)
	if err != nil {
		log.Printf("파일 내용 읽기 실패: %v", err)
		http.Error(w, "파일 내용 읽기 실패", http.StatusInternalServerError)
		return
	}

	log.Println("파일 읽기 성공")

	// JSON 형식의 개인 키 데이터를 PrivateKey 구조체로 변환
	var jsonPrivateKey PrivateKey
	err = json.Unmarshal(privateKeyData, &jsonPrivateKey)
	if err != nil {
		log.Printf("개인 키 JSON 파싱 실패: %v", err)
		http.Error(w, "개인 키 JSON 파싱 실패", http.StatusInternalServerError)
		return
	}

	// 타원 곡선 P-256 설정
	curve := elliptic.P256()
	privateKey := &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: curve,
			X:     jsonPrivateKey.X,
			Y:     jsonPrivateKey.Y,
		},
		D: jsonPrivateKey.D,
	}

	fmt.Printf("privateKey: %+v\n", privateKey)

	// // 서명할 메시지 설정
	// message := "hyeon" // vc
	// hash := sha256.Sum256([]byte(message))

	// // ECDSA 서명 생성
	// r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	// if err != nil {
	// 	log.Fatalf("서명 생성 실패: %v", err)
	// }

	// // 서명 결과를 바이트 배열로 결합
	// signature := append(r.Bytes(), s.Bytes()...)
	// fmt.Printf("서명: %x\n", signature)

	// err = os.WriteFile("Signature.txt", signature, 0644)
	// if err != nil {
	// 	log.Fatalf("서명 파일 저장 - 실패: %v", err)
	// }

	// fmt.Println("서명 파일 저장 - 성공")

	// 응답으로 성공 메시지 반환
	response := map[string]string{
		"message": "서명 생성이 완료되었습니다.",
		// "signatureFile": signatureFile,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
