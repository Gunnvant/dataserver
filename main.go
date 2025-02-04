package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Gunnvant/dataserver/handlers"
	"github.com/Gunnvant/dataserver/services"
)

func main() {
	creds := services.SqlCredentials{
		UserName: os.Getenv("MockDBUser"),
		Password: os.Getenv("MockDBPass"),
	}
	pgDb := services.PgDbConn{Credentials: creds, Host: "localhost", Port: "5432", DbName: "analytics"}
	Cnx := pgDb.Connect()

	stmt_handler := handlers.StatementHandler{Cnx: Cnx,
		AuthProvider: nil,
	}

	http.HandleFunc("/", stmt_handler.ServeHttp)
	fmt.Println("Server is running at port: 8080")
	http.ListenAndServe(":8080", nil)
}
