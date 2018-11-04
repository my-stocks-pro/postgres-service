package service


type Service struct {
	Conn
}

func New() Service {
	return Service{}
}

func InitDB(p Persist) (Conn, error) {
	conf, err := p.Init()
	if err != nil {
		panic(err)
	}

	conn, err := p.Client(conf)
	if err != nil {
		panic(err)
	}

	return conn, nil
}