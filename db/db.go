package db

import "tiny_twitter/customerrors"

type Storable interface {
	/*
	а что поанируется здесь возвращать по факту?
	*/
	GetData() interface{}
}

/*
зачем нужен этот интерфейс? а не просто []Storable или ...Storable
*/
type StorableList interface {
	At(i int) Storable
	Len() int
	Append(s Storable)
}

type updateFunction = func(s Storable) Storable

type QueryConstraints string

/*
Специфичные ошибки лучше держать в том же пакете
*/

type Database interface {
	Insert(s StorableList) customerrors.DbError
	Select(constraints QueryConstraints) (StorableList, customerrors.DbError)
	Update(constraints QueryConstraints, function updateFunction) customerrors.DbError
	Delete(constraints QueryConstraints) customerrors.DbError
	/*
	Знание о реализации лишнее, на то и интерфейс
	но если очент надо для каких-то целей, то лучше завести отдельный интерфейс
	*/
	GetDriverInfo() DriverConfig
}

type DriverConfig struct {
	// implementation-defined
}
