package impl

import (
	"sign_go/db"
	"sign_go/service"
	"sign_go/structs"
)

// KeyService 인터페이스 구현
type KeyServiceImpl struct {
	dao *db.MySQLDAO
}

// KeyServiceImpl 인스턴스 생성
func NewKeyServiceImpl(dao *db.MySQLDAO) service.KeyService {
	return &KeyServiceImpl{dao: dao}
}

// 개인키와 공개키 저장
func (s *KeyServiceImpl) SaveKeys(member structs.Member) error {
	return s.dao.SaveKeys(member)
}
