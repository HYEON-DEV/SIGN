package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" // MySQL 드라이버 - go get으로 설치 필요
)

// MySQLDAO 구조체 - 연동한 mysql DB 관리를 위한 구조체
type MySQLDAO struct {
	db *sql.DB
	// database/sql 패키지의 DB 타입을 포함하는 필드로, 실제 데이터베이스 연결을 나타낸다.
}

/*
MySQL 데이터베이스와 연결을 초기화, MySQLDAO 객체 생성

입력값
  - dsn(string) : 데이터 소스 이름 / MySQL DB에 연결하기 위한 정보를 포함한 문자열

출력값
  - MySQLDAO(구조체 포인터) : 생성된 mySQL 객체 포인터
  - error : 오류
*/
func NewMySQLDAO(dsn string) (*MySQLDAO, error) {

	// DB 연결 초기화 (열기)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("[ERROR] failed to initializing MySQL connection: %v\n", err)
		return nil, fmt.Errorf("failed to connect MySQL: %v", err)
	}

	// DB 연결 옵션 설정 - 이부분은 적당히 설정
	db.SetMaxOpenConns(100)                 // 최대 동시 연결 수
	db.SetMaxIdleConns(10)                  // 유휴 연결 수
	db.SetConnMaxLifetime(30 * time.Minute) // 연결 최대 생존시간

	// DB 연결 테스트 (없어도 됨)
	if err := db.Ping(); err != nil {
		db.Close()
		log.Printf("[ERROR] failed to test MySQL Ping: %v\n", err)
		return nil, fmt.Errorf("failed to ping MySQL: %v", err)
	}

	return &MySQLDAO{db: db}, nil
}

// mySQL 연결 종료
func (dao *MySQLDAO) Close() error {
	if err := dao.db.Close(); err != nil {
		return fmt.Errorf("failed to close MySQL connection: %v", err)
	}
	return nil
}
