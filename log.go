package log

type Logger interface {
	Tracef(string, ...interface{})
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
	Fatalf(string, ...interface{})
}

type DummyLogger bool

func NewDummyLogger() *DummyLogger {
	l := DummyLogger(false)

	return &l
}

func (l DummyLogger) Tracef(string, ...interface{}) {}
func (l DummyLogger) Debugf(string, ...interface{}) {}
func (l DummyLogger) Infof(string, ...interface{})  {}
func (l DummyLogger) Warnf(string, ...interface{})  {}
func (l DummyLogger) Errorf(string, ...interface{}) {}
func (l DummyLogger) Fatalf(string, ...interface{}) {}
