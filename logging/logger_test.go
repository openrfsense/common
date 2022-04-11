package logging

import (
	"io"
	"testing"
)

var testLogger = New(
	WithOutput(io.Discard),
	WithPrefix("benchmark"),
)

func TestNew(t *testing.T) {
	l := New(
		WithPrefix("test"),
		WithFlags(FlagsDevelopment),
		WithLevel(DebugLevel),
	)

	l.Debug("debug")
}

func BenchmarkLog(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testLogger.Info(b.N)
	}
}

func BenchmarkLog100(b *testing.B) {
	for i := 0; i < 100; i++ {
		for n := 0; n < b.N; n++ {
			testLogger.Info(b.N)
		}
	}
}

func BenchmarkLog1000(b *testing.B) {
	for i := 0; i < 10000; i++ {
		for n := 0; n < b.N; n++ {
			testLogger.Info(b.N)
		}
	}
}
