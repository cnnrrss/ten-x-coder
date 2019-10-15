package benchmarking

import "testing"

var x map[string]int

func BenchmarkEmptyMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x = make(map[string]int)
	}
}

type entry struct {
	v string
	c int
}

var e *entry

func BenchmarkEmptyEntry(b *testing.B) {
	for i := 0; i < b.N; i++ {
		e = &entry{}
	}
}
