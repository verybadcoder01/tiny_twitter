package server

import "tiny_twitter/db"

type SConfig struct {
	//implementation-defined
}

/*
Выглядит транспортом
тут задача - обрабатывать/преобразовывать запросы
зачем предоставлять из сервера бд?
также зачем конфиг? (он же нужен только при создании сервера)
*/

type Server interface {
	GetDB() db.Database
	GetConfig() SConfig
	/*
	очень не помешает вернуть error
	*/
	Handle(port string)
	Stop() error
}
