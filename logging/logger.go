package logging

type LogLevel = int

type Logger interface {
	Tracef(string)
	Debugf(string)
	Infof(string)
	Warnf(string)
	Errorf(string)
	Fatalf(string)
	SetLevel(level LogLevel)
}
