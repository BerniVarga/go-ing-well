package benchmarking

import (
	"fmt"
	"testing"
)

var gs string

func BenchmarkSprint(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		// have the function I'd like to benchmark here
		s = fmt.Sprint("Hello")
	}
	gs = s
}

func BenchmarkSprintfEmpty(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		// have the function I'd like to benchmark here
		s = fmt.Sprintf("Hello")
	}
	gs = s
}

func BenchmarkSprintfFormatted(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		// have the function I'd like to benchmark here
		s = fmt.Sprintf("Hello %d", i) // 8 byte
	}
	gs = s
}
