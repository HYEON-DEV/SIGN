package impl

import (
	"fmt"
	"log"
	"sign_go/db"
	"sign_go/service"
	"sign_go/structs"
	"time"
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

// SaveSession 메서드 구현
func (s *MemberServiceImpl) SaveSession(sessionID string, memberID int, name string, userID string, expiresAt time.Time) error {
	log.Println("SaveSession 호출, sessionID:", sessionID)
	if s.dao == nil {
		log.Println("DAO가 nil입니다.")
		return fmt.Errorf("DAO가 초기화되지 않았습니다.")
	}
	return s.dao.SaveSession(sessionID, memberID, name, userID, expiresAt)
}

// DeleteSession 메서드 구현
func (s *MemberServiceImpl) DeleteSession(sessionID string) error {
	log.Println("DeleteSession 호출, sessionID:", sessionID)
	if s.dao == nil {
		log.Println("DAO가 nil입니다.")
		return fmt.Errorf("DAO가 초기화되지 않았습니다.")
	}
	return s.dao.DeleteSession(sessionID)
}
