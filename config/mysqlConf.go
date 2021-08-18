package config

type Mysql struct {
	Path     string
	Username string
	Password string
	Dbname   string
	Config   string
}

func TestDsn() string {
	return ""
}

func DevDsn() string {
	return "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
}

func ProDsn() string {
	return ""
}
