package service

type Persist interface {
	Init() (Config, error)
	Client(Config) (Conn, error)
}









