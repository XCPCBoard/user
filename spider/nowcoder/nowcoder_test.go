package nowcoder

import (
	"fmt"
	"testing"
)

// @Author: Feng
// @Date: 2022/5/12 17:40

func BenchmarkAppend(b *testing.B) {
	foo := []string{}
	for i := 0; i < b.N; i++ {
		foo = append(foo, fmt.Sprintf("%v", i))
	}
}

func BenchmarkMap(b *testing.B) {
	foo := map[string]int{}
	for i := 0; i < b.N; i++ {
		foo[fmt.Sprintf("%v", i)] = i
	}
}
