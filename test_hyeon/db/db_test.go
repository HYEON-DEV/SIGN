package db

import (
	"database/sql"
	"testing"

	"test_hyeon/structs"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestSaveKeys(t *testing.T) {

	// sqlmock을 사용하여 데이터베이스 모킹
	// db: 모킹된 데이터베이스 연결
	// mock: 모킹 설정을 위한 핸들러
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
		PrivateKey: sql.NullString{String: `{"d":381574558907}`, Valid: true},
		PublicKey:  sql.NullString{String: `{"x":381574558907}`, Valid: true},
	}

	// 예상되는 쿼리와 결과 설정
	mock.ExpectExec(`UPDATE member
		SET private_key = \?, public_key = \?
		WHERE member_id = \?`).
		WithArgs(member.PrivateKey, member.PublicKey, member.MemberID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// SaveKeys 메서드 호출
	err = dao.SaveKeys(member)
	if err != nil {
		t.Errorf("키 업데이트 중 오류 발생: %v", err)
	}

	// 모든 예상 쿼리가 실행되었는지 확인
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("키 업데이트 쿼리 실행 오류: %v", err)
	}
}
