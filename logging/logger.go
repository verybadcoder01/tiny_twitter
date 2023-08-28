package logging

type LogLevel = int

/*
минивопрос - зачем суффикс *f? (тут же не формата)
мне очень нравится идея structured logging, так как добавляется гибкость при фильтрации записей в различных коллекторах логов
*/

type Logger interface {
	Tracef(string)
	Debugf(string)
	Infof(string)
	Warnf(string)
	Errorf(string)
	Fatalf(string)
	SetLevel(level LogLevel)
}
