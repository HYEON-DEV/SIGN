package db

type Member struct {
	MemberID   int    `db:"member_id"`
	Name       string `db:"name"`
	UserID     string `db:"user_id"`
	UserPW     string `db:"user_pw"`
	RegDate    string `db:"reg_date"`
	PrivateKey string `db:"private_key"`
	PublicKey  string `db:"public_key"`
	VC         string `db:"vc"`
	Facility   string `db:"facility"`
}
