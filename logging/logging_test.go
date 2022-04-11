package logging

import (
	"io"
	"testing"
)

func BenchmarkLog(b *testing.B) {
	logger.SetOutput(io.Discard)
	for n := 0; n < b.N; n++ {
		Info("benchmark", b.N)
	}
}

func BenchmarkLog100(b *testing.B) {
	logger.SetOutput(io.Discard)
	for i := 0; i < 100; i++ {
		for n := 0; n < b.N; n++ {
			Info("benchmark", b.N)
		}
	}
}

func BenchmarkLog1000(b *testing.B) {
	logger.SetOutput(io.Discard)
	for i := 0; i < 10000; i++ {
		for n := 0; n < b.N; n++ {
			Info("benchmark", b.N)
		}
	}
}
