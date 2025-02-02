package util

/*
 * 설치 - go get github.com/google/uuid
 * uuid 라이브러리를 사용하여 고유한 세션 ID를 생성,
 * GenerateSessionID 함수를 직접 구현하지 않아도 됨.
 */
import (
	"github.com/google/uuid"
)

// GenerateSessionID 함수는 새로운 UUID를 생성하여 세션 ID로 반환합니다.
func GenerateSessionID() string {
	return uuid.New().String()
}
