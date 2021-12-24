package inter

import "math/big"

type EthApiBackend struct {
}

func (b *EthApiBackend) SyncProgress() int {
	return 1
}

func (b *EthApiBackend) FeeHistory() big.Int {
	return *big.NewInt(1)
}
