package db

import "tiny_twitter/customerrors"

type Storable interface {
	GetData() interface{}
}

type StorableList interface {
	At(i int) Storable
	Len() int
	Append(s Storable)
}

type updateFunction = func(s Storable) Storable

type QueryConstraints string

type Database interface {
	Insert(s StorableList) customerrors.DbError
	Select(constraints QueryConstraints) (StorableList, customerrors.DbError)
	Update(constraints QueryConstraints, function updateFunction) customerrors.DbError
	Delete(constraints QueryConstraints) customerrors.DbError
	GetDriverInfo() DriverConfig
}

type DriverConfig struct {
	// implementation-defined
}
