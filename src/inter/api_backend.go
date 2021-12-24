package inter

type EthApiBackend struct {
}

func (b *EthApiBackend) SyncProgress() int {
	return 1
}
