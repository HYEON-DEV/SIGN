/*
 * 서비스 인터페이스
 * 비즈니스 프로세스 처리
 * 사용할 repository 결정
 * repository 기능 호출해서 데이터 처리
 */

package service

import (
	"sign_go/structs"
)

/*
 * KeyService 인터페이스 - 비즈니스 로직 정의
 */
type KeyService interface {
	// 생성한 키 DB에 저장하는 함수
	SaveKeys(member structs.Member) error
	// 키 존재 여부 확인하는 함수
	CheckKeys(memberID int) (bool, error)
}

/*
 * interface
 *
 * 메서드의 집합을 정의하는 타입
 */
