package services

import (
	"database/sql"
	"dataserver/entities"
	"log"

	_ "github.com/lib/pq"
)

// DB interface

// App db credentials
type SqlCredentials struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// Postgress Object
type PgDbConn struct {
	Credentials SqlCredentials `json:"credentials"`
	Host        string         `json:"host"`
	Port        string         `json:"port"`
	DbName      string         `json:"DbName"`
}
type SqliteConn struct {
	Path string
}

type SqlServerConn struct {
	Credentials SqlCredentials `json:"credentials"`
	Host        string         `json:"host"`
	Port        string         `json:"port"`
	DbName      string         `json:"DbName"`
}

// Connnect to postgres database
func (p PgDbConn) Connect() *entities.Cnx {
	url := "postgresql://" + p.Credentials.UserName + ":" + p.Credentials.Password + "@" + p.Host + ":" + p.Port + "/" + p.DbName + "?sslmode=disable"
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Panicf("Error while connecting with db %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Panicf("Error connecting to database %v", err)
	}
	log.Printf("Connected successfully to db %v", p.DbName)
	return &entities.Cnx{DB: db, Type: "pg"}
}

// Connect to Sql Server

func (p SqlServerConn) Connect() *entities.Cnx {
	panic("Not implemented")
	//return &entities.Cnx{DB: nil, Type: "sqlserver"}
}

// Connect sqlite
func (p SqliteConn) Connect() *entities.Cnx {
	panic("Not implemented")
	//return &entities.Cnx{DB: nil, Type: ""}
}
