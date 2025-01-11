package services

import (
	"log"

	"github.com/Gunnvant/dataserver/entities"
)

func GetDataFromQuery(cnx *entities.Cnx, q string) (entities.SqlResponse, error) {
	db := cnx.DB
	rows, err := db.Query(q)
	if err != nil {
		log.Printf("Couldn't execute the query. Error %v", err)
		return entities.SqlResponse{}, err
	}
	defer rows.Close()
	var result entities.SqlResponse
	columns, err := rows.Columns()
	if err != nil {
		log.Printf("Can't get column names. Error: %v", err)
		return entities.SqlResponse{}, err
	}
	values := make([]interface{}, len(columns))
	for rows.Next() {
		// Scan the row into the values slice
		for i := range values {
			values[i] = new(interface{})
		}

		// Scan the row into the values slice
		err := rows.Scan(values...)
		if err != nil {
			log.Printf("Unable to scan row. Error: %v", err)
		}
		rowMap := make(map[string]interface{})
		for i, colName := range columns {
			// Dereference the pointer to get the actual value
			rowMap[colName] = *(values[i].(*interface{}))
		}
		result.Resp = append(result.Resp, rowMap)
	}
	return result, nil
}
