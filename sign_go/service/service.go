/*
 * 서비스 인터페이스
 * 비즈니스 프로세스 처리
 * 사용할 repository 결정
 * repository 기능 호출해서 데이터 처리
 */

package service

import (
	"test_hyeon/structs"
)

/*
KeyService 인터페이스 - 비즈니스 로직 정의
*/
type KeyService interface {
	SaveKeys(member structs.Member) error
}
