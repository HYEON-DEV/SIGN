package structs

import "math/big"

// PrivateKey 구조체 정의
type PrivateKey struct {
	D     *big.Int `json:"d"`
	X     *big.Int `json:"x"`
	Y     *big.Int `json:"y"`
	Curve string   `json:"curve"`
}

// PublicKey 구조체 정의
type PublicKey struct {
	X     *big.Int `json:"x"`
	Y     *big.Int `json:"y"`
	Curve string   `json:"curve"`
}
