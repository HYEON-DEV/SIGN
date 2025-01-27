package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

// MySQLDAO 구조체 - 연동한 mysql DB 관리를 위한 구조체
type MySQLDAO struct {
	db *sql.DB
}

// mySQL 연결 초기화 및 객체 생성
//
// 입력값
//   - dsn(string) : mySQL DB 연결에 필요한 정보를 포함한 문자열
//
// 출력값
//   - MySQLDAO(구조체 포인터) : 생성된 mySQL 객체 포인터
//   - error : 오류
func NewMySQLDAO(dsn string) (*MySQLDAO, error) {
	// DB 연결 초기화(열기)
	db, err := sql.Open("mysql", dsn)
	if err != nil { // 오류처리
		log.Printf("[ERROR] failed to initializing MySQL connection: %v\n", err)
		return nil, fmt.Errorf("failed to connect MySQL: %v", err)
	}

	// DB 연결 옵션 설정 - 이부분은 적당히 설정하면 됨
	db.SetMaxOpenConns(100)                 // 최대 동시 연결 수
	db.SetMaxIdleConns(10)                  // 유후 연결 수
	db.SetConnMaxLifetime(30 * time.Minute) // 연결 최대 생존시간

	// 연결 테스트(ping) - 없어도 되는 부분
	if err := db.Ping(); err != nil {
		db.Close() // 연결 테스트 실패시 DB 닫기
		log.Printf("[ERROR] failed to test MySQL Ping: %v\n", err)
		return nil, fmt.Errorf("failed to ping MySQL: %v", err)
	}

	return &MySQLDAO{db: db}, nil
}

// mySQL 연결 종료 함수
func (dao *MySQLDAO) Close() error {
	if err := dao.db.Close(); err != nil {
		return fmt.Errorf("failed to close MySQL connection: %v", err)
	}
	return nil
}
