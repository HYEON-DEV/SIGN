package db

import (
	"encoding/json"
	"testing"

	"sign_go/structs"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestMemberLogin(t *testing.T) {
	// sqlmock을 사용하여 데이터베이스 모킹
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("모킹된 데이터베이스 연결 오류 : %v", err)
	}
	defer db.Close()

	// MySQLDAO 인스턴스 생성, 모킹된 데이터베이스 연결 설정
	dao := &MySQLDAO{db: db}

	// 테스트할 member 데이터 생성
	member := structs.Member{
		MemberID:   1,
		Name:       "Test User",
		UserID:     "testuser",
		UserPW:     "123qwe!@#",
		PrivateKey: json.RawMessage(`{"d":381574558907}`),
		PublicKey:  json.RawMessage(`{"x":381574558907}`),
	}

	// 예상되는 쿼리와 결과 설정
	rows := sqlmock.NewRows([]string{"member_id", "name", "user_id", "user_pw", "reg_date", "private_key", "public_key", "vc", "facility"}).
		AddRow(member.MemberID, member.Name, member.UserID, member.UserPW, member.RegDate, member.PrivateKey, member.PublicKey, member.VC, member.Facility)

	mock.ExpectQuery(`SELECT member_id, name, user_id, user_pw, reg_date, private_key, public_key, vc, facility FROM member WHERE user_id = \? AND user_pw = \?`).
		WithArgs(member.UserID, member.UserPW).
		WillReturnRows(rows)

	// MemberLogin 메서드 호출
	result, err := dao.MemberLogin(member.UserID, member.UserPW)
	if err != nil {
		t.Errorf("로그인 중 오류 발생: %v", err)
	}

	// 결과 검증
	if result == nil || result.MemberID != member.MemberID {
		t.Errorf("예상된 멤버와 일치하지 않음: %v", result)
	}

	// 모든 예상 쿼리가 실행되었는지 확인
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("로그인 쿼리 실행 오류: %v", err)
	}
}
