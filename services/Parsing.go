package services

import (
	"dataserver/entities"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateFilterOps(s []string) error {
	valid_ops := map[string]bool{
		"=":   true,
		"<=":  true,
		">=":  true,
		"<":   true,
		">":   true,
		"<>":  true,
		"IN":  true,
		"OR":  true,
		"AND": true,
		"NOT": true,
	}
	for _, val := range s {
		if !valid_ops[strings.ToUpper(val)] {
			return errors.New("Invalid operator: " + val)
		}
	}
	return nil
}
func CreateSqlParams(r *http.Request) (entities.SqlParams, error) {
	var sql_params entities.SqlParams
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Unable to read the body: %v", err)
		return entities.SqlParams{}, err
	}
	defer r.Body.Close()
	err = json.Unmarshal(body, &sql_params)
	if err != nil {
		log.Printf("Can't decode the body sent for sql parameters: %v", err)
		return entities.SqlParams{}, err
	}

	validate := validator.New()
	err = validate.Struct(sql_params)
	if err != nil {
		log.Printf("Validation failed on the sql params with error: %v", err)
		return entities.SqlParams{}, err
	}
	if sql_params.Filters != nil {
		err = ValidateFilterOps(sql_params.Filters.FilterOps)
		if err != nil {
			log.Printf("Validation failure for filter operations: %v", err)
			return entities.SqlParams{}, err
		}
	}
	return sql_params, nil
}

func AddWhereClause(s entities.SqlParams, sql_stmt string) string {
	sql_stmt += ` WHERE `
	for i := 0; i < len(s.Filters.FilterCols); i++ {
		sql_stmt += s.Filters.FilterCols[i] + " " + s.Filters.FilterOps[i] + " " + s.Filters.FilterVals[i] + " "
	}
	return sql_stmt
}

func AddSortClause(s entities.SqlParams, sql_stmt string) string {
	sql_stmt += `ORDER BY `
	for i := 0; i < len(s.Sort.SortCols); i++ {
		if i < len(s.Sort.SortCols)-1 {
			sql_stmt += s.Sort.SortCols[i] + " " + s.Sort.SortType[i] + ", "
		} else {
			sql_stmt += s.Sort.SortCols[i] + " " + s.Sort.SortType[i] + " "
		}
	}
	return sql_stmt
}

func CreateSqlStatementPg(s entities.SqlParams) string {
	var sql_stmt string
	if *s.Distinct {
		sql_stmt = `Select Distinct ` + strings.Join(s.SelectCols, ",") + ` from ` + s.Table
	} else {
		sql_stmt = `Select ` + strings.Join(s.SelectCols, ",") + ` from ` + s.Table
	}
	if s.Filters != nil {
		sql_stmt = AddWhereClause(s, sql_stmt)
	}
	if s.Sort != nil {
		sql_stmt = AddSortClause(s, sql_stmt)
	}
	if s.LimitVals != nil {
		sql_stmt += `Limit ` + *s.LimitVals
	}
	sql_stmt = strings.TrimSpace(sql_stmt)
	log.Printf("The sql query is: %s", sql_stmt)
	return sql_stmt
}

func CreateSqlStatementSqlite(s entities.SqlParams) string {
	var sql_stmt string
	if *s.Distinct {
		sql_stmt = `Select Distinct ` + strings.Join(s.SelectCols, ",") + ` from ` + s.Table
	} else {
		sql_stmt = `Select ` + strings.Join(s.SelectCols, ",") + ` from ` + s.Table
	}
	if s.Filters != nil {
		sql_stmt = AddWhereClause(s, sql_stmt)
	}

	if s.Sort != nil {
		sql_stmt = AddSortClause(s, sql_stmt)
	}

	if s.LimitVals != nil {
		sql_stmt += `Limit ` + *s.LimitVals
	}
	sql_stmt = strings.TrimSpace(sql_stmt)
	log.Printf("The sql query is: %s", sql_stmt)
	return sql_stmt
}

func CreateSqlStatementSqlServer(s entities.SqlParams) string {
	var sql_stmt string
	if *s.Distinct {
		if s.LimitVals != nil {
			sql_stmt = `Select Distinct TOP ` + *s.LimitVals + " " + strings.Join(s.SelectCols, ",") + ` from ` + s.Table
		} else {
			sql_stmt = `Select Distinct ` + strings.Join(s.SelectCols, ",") + ` from ` + s.Table
		}
	} else {
		if s.LimitVals != nil {
			sql_stmt = `Select TOP ` + *s.LimitVals + " " + strings.Join(s.SelectCols, ",") + ` from ` + s.Table
		} else {
			sql_stmt = `Select ` + strings.Join(s.SelectCols, ",") + ` from ` + s.Table
		}
	}

	if s.Filters != nil {
		sql_stmt = AddWhereClause(s, sql_stmt)
	}
	if s.Sort != nil {
		sql_stmt = AddSortClause(s, sql_stmt)
	}
	sql_stmt = strings.TrimSpace(sql_stmt)
	log.Printf("The sql query is: %s", sql_stmt)
	return sql_stmt
}
func GetStatementSql(dialect *string, sql_params entities.SqlParams) string {
	switch *dialect {
	case "pg":
		return CreateSqlStatementPg(sql_params)
	case "sqlserver":
		return CreateSqlStatementSqlServer(sql_params)
	default:
		return CreateSqlStatementSqlite(sql_params)
	}
}
