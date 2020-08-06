package repeat

import "testing"

// go 中的测试文件应该命名为name_test，否则找不到
// https://stackoverflow.com/questions/28240489/golang-testing-no-test-files
// go test
func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"
	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}

// Alt + <- 返回上一个光标位置
// windows go test -bench="."
// others go test -bench=.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
