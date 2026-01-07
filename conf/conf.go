package conf

type Config struct {
	Username string
	Password string
	Data     *data
}

type data struct {
	Dsn string
}
