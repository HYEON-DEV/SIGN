package main_ex

import (
	"database/sql"
	"fmt"
	"log"
	"test_hyeon/config"
	"test_hyeon/db"
	"test_hyeon/util"
)

func main_ex() {
	util.LogSetup()
	util.Enterlog("main")
	defer util.Leavelog("main")

	// Load configuration
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database
	dsn := config.MySQLDSN()
	mysqlDAO, err := db.NewMySQLDAO(dsn) // db 패키지에서 DB를 실제로 연동하는 함수 호출
	if err != nil {
		log.Fatalf("[ERROR] Failed to connect MySQL: %v", err)
	}
	defer func() { // 프로그램 종료 시 DB를 닫도록 설정
		if mysqlDAO != nil {
			if err := mysqlDAO.Close(); err != nil {
				log.Printf("[WARN] Failed to close mySQL connection: %v", err)
			}
		}
	}()

	// Test: Get all members
	members, err := mysqlDAO.GetAllMembers()
	if err != nil {
		util.Errlog("main", "DB_ERR", "Error getting all members", err)
		log.Fatalf("Error getting all members: %v", err)
	}

	// Print all members
	for _, member := range members {
		fmt.Printf("ID: %d, Name: %s, UserID: %s, UserPW: %s, RegDate: %s, PrivateKey: %s, PublicKey: %s, VC: %s, Facility: %s\n",
			member.MemberID, member.Name, member.UserID, member.UserPW, member.RegDate, nullStringToString(member.PrivateKey), nullStringToString(member.PublicKey), nullStringToString(member.VC), nullStringToString(member.Facility))
	}
}

func nullStringToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return "NULL"
}
