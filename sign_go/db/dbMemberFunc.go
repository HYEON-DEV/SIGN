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
)

// MySQLDAO 구조체 - 데이터베이스 연결 관리
// type MySQLDAO struct {
//     db *sql.DB
// }

func (dao *MySQLDAO) MemberLogin(user_id string, user_pw string) (*structs.Member, error) {
	util.Enterlog("MemberLogin")
	defer util.Leavelog("MemberLogin")

	query := `
		SELECT member_id, name, user_id, user_pw, reg_date, 
				private_key, public_key, vc, facility 
		FROM member
		WHERE user_id = ? AND user_pw = ?`

	log.Println("[SELECT] id,pw로 로그인")
	row := dao.db.QueryRow(query, user_id, user_pw)

	var member structs.Member
	err := row.Scan(&member.MemberID, &member.Name, &member.UserID, &member.UserPW, &member.RegDate, &member.PrivateKey, &member.PublicKey, &member.VC, &member.Facility)

	if err != nil {
		return nil, fmt.Errorf("로그인 실패: %v", err)
	}

	return &member, nil
}
