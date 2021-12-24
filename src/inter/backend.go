package inter

import "math/big"

type Backend interface {
	SyncProgress() int

	FeeHistory() big.Int
}
