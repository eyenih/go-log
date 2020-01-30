package log

import (
	"testing"
)

func TestDummyLogger(t *testing.T) {
	l := NewDummyLogger()

	func(l Logger) {
		l.Tracef("%d", 1)
		l.Debugf("%d", 1)
		l.Infof("%d", 1)
		l.Warnf("%d", 1)
		l.Errorf("%d", 1)
		l.Fatalf("%d", 1)
	}(l)
}

func BenchmarkDummyLogger(b *testing.B) {
	l := NewDummyLogger()

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		l.Tracef("%d", 1)
		l.Debugf("%d", 1)
		l.Infof("%d", 1)
		l.Warnf("%d", 1)
		l.Errorf("%d", 1)
		l.Fatalf("%d", 1)
	}
}
