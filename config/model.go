package config

type ConfigStruct struct {
	Server Server
	Db     Db
}

type Server struct {
	Host string
	Port string
}

type Db struct {
	Host     string
	Port     string
	User     string
	Password string
}
