package log

import (
	"testing"
)

func TestDummyLogger(t *testing.T) {
	l := NewDummyLogger()

	func(l Logger) {
		l.Trace("", make(map[string]interface{}, 1))
		l.Debug("", make(map[string]interface{}, 1))
		l.Info("", make(map[string]interface{}, 1))
		l.Warn("", make(map[string]interface{}, 1))
		l.Error("", make(map[string]interface{}, 1))
		l.Fatal("", make(map[string]interface{}, 1))
	}(l)
}
