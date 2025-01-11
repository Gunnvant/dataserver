package entities

import "database/sql"

type Cnx struct {
	DB   *sql.DB
	Type string
}

type Config struct {
	Port string `json:"port"`
	Host string `json:"host"`
	User string `json:"user"`
	Pwd  string `json:"pwd"`
	Db   string `json:"db"`
}
