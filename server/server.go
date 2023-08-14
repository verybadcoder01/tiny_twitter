package server

import "tiny_twitter/db"

type SConfig struct {
	//implementation-defined
}

type Server interface {
	GetDB() db.Database
	GetConfig() SConfig
	Handle(port string)
	Stop() error
}
