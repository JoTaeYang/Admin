package gen

import "testing"

func BenchmarkUUID(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		UUID()
	}
}

func BenchmarkSnowFlake(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		SnowFlake()
	}
}
