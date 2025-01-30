package main

import (
	"database/sql"
	"fmt"
	"log"
	"test_hyeon/config"
	"test_hyeon/db"
	"test_hyeon/util"
)

func main() {
	util.LogSetup()
	util.Enterlog("main")
	defer util.Leavelog("main")

	// Load configuration
	_, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Initialize database
	dao, err := db.NewMySQLDAO(config.GConfig.MySQLDSN())
	if err != nil {
		util.Errlog("main", "DB_ERR", "Error initializing database", err)
		log.Fatalf("Error initializing database: %v", err)
	}
	defer dao.Close()

	// Test: Get all members
	members, err := dao.GetAllMembers()
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
