package db

import (
	"fmt"
	"test_hyeon/util"
)

// member_id로 member 조회
func (dao *MySQLDAO) GetMemberByID(id int) (*Member, error) {
	util.Enterlog("GetMemberByID")
	defer util.Leavelog("GetMemberByID")

	query := `
		SELECT member_id, name, user_id, user_pw, reg_date, 
				private_key, public_key, vc, facility 
		FROM member
		WHERE member_id = ?`

	row := dao.db.QueryRow(query, id)

	var member Member
	err := row.Scan(&member.MemberID, &member.Name, &member.UserID, &member.UserPW, &member.RegDate, &member.PrivateKey, &member.PublicKey, &member.VC, &member.Facility)

	if err != nil {
		return nil, fmt.Errorf("failed to get member by ID: %v", err)
	}

	return &member, nil
}

// 모든 member 조회
func (dao *MySQLDAO) GetAllMembers() ([]Member, error) {
	util.Enterlog("GetAllMembers")
	defer util.Leavelog("GetAllMembers")

	query := `
		SELECT member_id, name, user_id, user_pw, reg_date, 
				private_key, public_key, vc, facility 
		FROM member`

	rows, err := dao.db.Query(query)

	if err != nil {
		return nil, fmt.Errorf("failed to get all members: %v", err)
	}
	// 반복이 끝나면 rows.Close()를 호출하여 리소스를 해제
	defer rows.Close()

	var members []Member

	for rows.Next() {

		var member Member

		// rows.Scan()을 사용하여 Member 구조체에 매핑
		err := rows.Scan(&member.MemberID, &member.Name, &member.UserID, &member.UserPW, &member.RegDate, &member.PrivateKey, &member.PublicKey, &member.VC, &member.Facility)

		if err != nil {
			return nil, fmt.Errorf("failed to scan member: %v", err)
		}

		// 매핑된 Member 구조체를 members 슬라이스에 추가
		members = append(members, member)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %v", err)
	}

	return members, nil
}
