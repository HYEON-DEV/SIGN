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
	"test_hyeon/structs"
	"test_hyeon/util"
)

// MySQLDAO 구조체 - 데이터베이스 연결 관리
// type MySQLDAO struct {
//     db *sql.DB
// }

/*
개인키, 공개키 저장
*/
func (dao *MySQLDAO) SaveKeys(member structs.Member) error {
	util.Enterlog("SaveKeys")
	defer util.Leavelog("SaveKeys")

	query := `
		UPDATE member
		SET private_key = ?, public_key = ?
		WHERE member_id = ?	`

	log.Println("[UPDATE] DB에 개인·공개키 저장")

	// sql.Result 타입으로 반환
	result, err := dao.db.Exec(query, member.PrivateKey, member.PublicKey, member.MemberID)
	if err != nil {
		log.Printf("DB에 키 저장 실패: %v\n", err)
		return err
	}

	row, err := result.RowsAffected()
	if err != nil {
		log.Printf("RowsAffected 오류: %v\n", err)
		return err
	}
	log.Println("update count: ", row)

	return nil
}

// member_id로 member 조회
func (dao *MySQLDAO) GetMemberByID(id int) (*structs.Member, error) {
	util.Enterlog("GetMemberByID")
	defer util.Leavelog("GetMemberByID")

	query := `
		SELECT member_id, name, user_id, user_pw, reg_date, 
				private_key, public_key, vc, facility 
		FROM member
		WHERE member_id = ?`

	row := dao.db.QueryRow(query, id)

	var member structs.Member
	err := row.Scan(&member.MemberID, &member.Name, &member.UserID, &member.UserPW, &member.RegDate, &member.PrivateKey, &member.PublicKey, &member.VC, &member.Facility)

	if err != nil {
		return nil, fmt.Errorf("failed to get member by ID: %v", err)
	}

	return &member, nil
}

/*
모든 member 조회
*/
func (dao *MySQLDAO) GetAllMembers() ([]structs.Member, error) {
	util.Enterlog("GetAllMembers")
	defer util.Leavelog("GetAllMembers")

	query := `
		SELECT member_id, name, user_id, user_pw, reg_date, 
				private_key, public_key, vc, facility 
		FROM member`

	log.Println("[SELECT] member 데이터 전체 조회")
	rows, err := dao.db.Query(query)

	if err != nil {
		util.Errlog("GetAllMembers", "DB_ERR", "Failed to get all members", err)
		return nil, fmt.Errorf("failed to get all members: %v", err)
	}
	// 반복이 끝나면 rows.Close()를 호출하여 리소스 해제
	defer rows.Close()

	var members []structs.Member

	for rows.Next() {

		var member structs.Member

		// rows.Scan()을 사용하여 Member 구조체에 매핑
		err := rows.Scan(&member.MemberID, &member.Name, &member.UserID, &member.UserPW, &member.RegDate, &member.PrivateKey, &member.PublicKey, &member.VC, &member.Facility)

		if err != nil {
			util.Errlog("GetAllMembers", "DB_ERR", "Failed to scan member", err)
			return nil, fmt.Errorf("failed to scan member: %v", err)
		}

		// 매핑된 Member 구조체를 members 슬라이스에 추가
		members = append(members, member)
	}

	if err = rows.Err(); err != nil {
		util.Errlog("GetAllMembers", "DB_ERR", "Rows iteration error", err)
		return nil, fmt.Errorf("rows iteration error: %v", err)
	}

	return members, nil
}
