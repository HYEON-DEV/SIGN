/*
 * 모델
 */

package structs

import "encoding/json"

type Member struct {
	MemberID   int             `db:"member_id"`
	Name       string          `db:"name"`
	UserID     string          `db:"user_id"`
	UserPW     string          `db:"user_pw"`
	RegDate    string          `db:"reg_date"`
	PrivateKey json.RawMessage `db:"private_key"`
	PublicKey  json.RawMessage `db:"public_key"`
	VC         json.RawMessage `db:"vc"`
	Facility   json.RawMessage `db:"facility"`
}
