package atcoder

import (
	"XCPCer_board/model"
	"testing"
)

func BenchmarkFlush(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Flush(model.TestAtcIdLQY)
	}
}
