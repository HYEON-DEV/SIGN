package db

import "database/sql"

type Member struct {
	MemberID   int            `db:"member_id"`
	Name       string         `db:"name"`
	UserID     string         `db:"user_id"`
	UserPW     string         `db:"user_pw"`
	RegDate    string         `db:"reg_date"`
	PrivateKey sql.NullString `db:"private_key"`
	PublicKey  sql.NullString `db:"public_key"`
	VC         sql.NullString `db:"vc"`
	Facility   sql.NullString `db:"facility"`
}
