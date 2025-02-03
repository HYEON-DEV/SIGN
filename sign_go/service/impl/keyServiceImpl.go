package impl

import (
	"sign_go/db"
	"sign_go/service"
	"sign_go/structs"
)

// KeyService 인터페이스 구현
type KeyServiceImpl struct {
	// DB 접근 객체를 사용해 DB와 상호작용한다.
	// dbSet.go에서 정의 - DB 연결 관리, DB 작업 수행하는 메서드를 포함한다.
	dao *db.MySQLDAO

	// KeyServiceImpl 구조체는 KeyService 인터페이스를 구현하며,
	//    DB 작업을 수행하기 위해 MySQLDAO 객체를 사용한다.
}

// KeyServiceImpl 인스턴스 생성
func NewKeyServiceImpl(dao *db.MySQLDAO) service.KeyService {
	return &KeyServiceImpl{dao: dao}
}

// 개인키와 공개키 저장
func (s *KeyServiceImpl) SaveKeys(member structs.Member) error {
	return s.dao.SaveKeys(member)
	// dao 필드를 사용해 키를 DB에 저장한다.
}

// 키 존재 유무 확인
func (s *KeyServiceImpl) CheckKeys(memberID int) (bool, error) {
	return s.dao.CheckKeys(memberID)
	// dao 필드를 사용해 키 존재 여부를 확인한다.
}
