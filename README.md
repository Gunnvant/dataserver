## Data Server

**Purpose:**

Expose a pg sql table via API.

**How it works:**

You need to generate a `main.go` file. Supply the details of your database, whose tables you want to expose.

```go
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

```
You can then send a post request that adheres to the following struct:

```go
type SqlParams struct {
	SelectCols []string      `json:"select_cols" validate:"min=1"`
	Table      string        `json:"table" validate:"required"`
	LimitVals  *string       `json:"limit_vals,omitempty"`
	Filters    *FilterParams `json:"filter_params,omitempty"`
	Sort       *SortParams   `json:"sort_params,omitempty"`
	Distinct   *bool         `json:"distinct"`
}

```

You will then get a response in json format.

**Example request:**

```json
{
  "select_cols":["*"],
  "table":"roles",
  "filter_params":{
    "filter_cols":["user_name"],
    "filter_ops":["="],
    "filter_vals":["'gunnvant3.14@gmail.com'"]
  },
  "distinct":false
}

```

**Example Response**

```json
{
  "response": [
    {
      "id": 1,
      "role": "admin",
      "user_name": "gunnvant3.14@gmail.com"
    },
    {
      "id": 2,
      "role": "user",
      "user_name": "gunnvant3.14@gmail.com"
    }
  ]
}
```
