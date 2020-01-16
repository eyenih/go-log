package log

type Logger interface {
	Trace(string, map[string]interface{})
	Debug(string, map[string]interface{})
	Info(string, map[string]interface{})
	Warn(string, map[string]interface{})
	Error(string, map[string]interface{})
	Fatal(string, map[string]interface{})
}

type DummyLogger bool

func NewDummyLogger() *DummyLogger {
	l := DummyLogger(false)

	return &l
}

func (l DummyLogger) Trace(string, map[string]interface{}) {}
func (l DummyLogger) Debug(string, map[string]interface{}) {}
func (l DummyLogger) Info(string, map[string]interface{})  {}
func (l DummyLogger) Warn(string, map[string]interface{})  {}
func (l DummyLogger) Error(string, map[string]interface{}) {}
func (l DummyLogger) Fatal(string, map[string]interface{}) {}
