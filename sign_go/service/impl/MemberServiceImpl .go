package impl

import (
	"fmt"
	"log"
	"sign_go/db"
	"sign_go/service"
	"sign_go/structs"
)

// MemberService 인터페이스 구현
type MemberServiceImpl struct {
	dao *db.MySQLDAO
}

// MemberServiceImpl 인스턴스 생성
func NewMemberServiceImpl(dao *db.MySQLDAO) (service.MemberService, error) {
	if dao == nil {
		return nil, fmt.Errorf("dao는 nil일 수 없습니다")
	}
	return &MemberServiceImpl{dao: dao}, nil
}

// MemberLogin 메서드 구현
func (s *MemberServiceImpl) MemberLogin(userID string, userPW string) (*structs.Member, error) {
	log.Println("MemberLogin 호출, userID:", userID)
	if s.dao == nil {
		log.Println("DAO가 nil입니다.")
		return nil, fmt.Errorf("DAO가 초기화되지 않았습니다.")
	}
	return s.dao.MemberLogin(userID, userPW)
}
