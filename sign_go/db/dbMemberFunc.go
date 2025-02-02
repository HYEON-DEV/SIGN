/*
 * Mapper역할
 * CRUD
 * DB의 일반적인 기능 담당
 * Repository (데이터 관리) (비즈니스로직과 DB 분리)
 */

package db

import (
	"fmt"
	"log"
	"sign_go/structs"
	"sign_go/util"
	"time"
)

func (dao *MySQLDAO) MemberLogin(user_id string, user_pw string) (*structs.Member, error) {
	util.Enterlog("MemberLogin")
	defer util.Leavelog("MemberLogin")

	query := `
	SELECT member_id, name, user_id, user_pw 
	FROM member
	WHERE user_id = ? AND user_pw = ?`

	log.Println("[SELECT] id,pw로 로그인")
	row := dao.db.QueryRow(query, user_id, user_pw)

	var member structs.Member
	// err := row.Scan(&member.MemberID, &member.Name, &member.UserID, &member.UserPW, &member.RegDate, &member.PrivateKey, &member.PublicKey, &member.VC, &member.Facility)
	err := row.Scan(&member.MemberID, &member.Name, &member.UserID, &member.UserPW)

	if err != nil {
		return nil, fmt.Errorf("로그인 실패: %v", err)
	}

	return &member, nil
}

// SaveSession - 세션 데이터를 DB에 저장
func (dao *MySQLDAO) SaveSession(sessionID string, memberID int, name string, userID string, expiresAt time.Time) error {
	util.Enterlog("SaveSession")
	defer util.Leavelog("SaveSession")

	query := `
        INSERT INTO session (session_id, member_id, name, user_id, expires_at)
        VALUES (?, ?, ?, ?, ?)`
	_, err := dao.db.Exec(query, sessionID, memberID, name, userID, expiresAt)
	if err != nil {
		util.Errlog("SaveSession", "DB_ERROR", "세션 저장 실패", err)
		return fmt.Errorf("세션 저장 실패: %v", err)
	}

	util.Enterlog(fmt.Sprintf("세션 저장 성공: sessionID=%s, memberID=%d, name=%s, userID=%s, expiresAt=%s", sessionID, memberID, name, userID, expiresAt))
	return nil
}

// DeleteSession - 세션 데이터를 DB에서 삭제
func (dao *MySQLDAO) DeleteSession(sessionID string) error {
	util.Enterlog("DeleteSession")
	defer util.Leavelog("DeleteSession")

	query := `DELETE FROM session WHERE session_id = ?`
	_, err := dao.db.Exec(query, sessionID)
	if err != nil {
		util.Errlog("DeleteSession", "DB_ERROR", "세션 삭제 실패", err)
		return fmt.Errorf("세션 삭제 실패: %v", err)
	}

	util.Enterlog(fmt.Sprintf("세션 삭제 성공: sessionID=%s", sessionID))
	return nil
}
