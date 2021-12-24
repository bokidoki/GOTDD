package inter

import (
	"fmt"
	"testing"
)

func TestBackend(t *testing.T) {
	api := &EthApiBackend{}
	API(api)
}

func API(backend Backend) {
	fmt.Println(backend.FeeHistory())
}
