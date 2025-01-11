package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Gunnvant/dataserver/entities"
	"github.com/Gunnvant/dataserver/services"

	"github.com/stretchr/testify/assert"
)

func TestValidateFilterOps(t *testing.T) {
	valid_ops := []string{"=", "<=",
		">=", "<",
		">", "<>",
		"IN", "OR",
		"AND", "NOT"}
	invalid_ops := []string{"=", "<=",
		">=", "<",
		">", "<>",
		"IN", "OR",
		"AND", "NOTY"}
	err1 := services.ValidateFilterOps(valid_ops)
	err2 := services.ValidateFilterOps(invalid_ops)
	assert.Equal(t, err1, nil)
	assert.Error(t, err2)
}

func TestSqlParamCreator(t *testing.T) {
	simpleBody := `{
					"select_cols":["Col1","Col2"],
					"table":"t1",
					"distinct":true
				}`
	bodyWithFilter := `{
					"select_cols":["Col1","Col2"],
					"table":"t1",
					"distinct":true,
					"filter_params":{
						"filter_cols":["C3","and C4"],
						"filter_vals":["V1","V2"],
						"filter_ops":["=","<="]
					}
	}`
	bodyFailWithOpsValidError := `{
			"select_cols":["Col1","Col2"],
					"table":"t1",
					"distinct":true,
					"filter_params":{
						"filter_cols":["C3","and C4"],
						"filter_vals":["V1","V2"],
						"filter_ops":["=","!="]
					}
	}`
	bodyFailWithValidError := `{
		"select_cols":[],
				"table":"t1",
				"distinct":true,
				"filter_params":{
					"filter_cols":["C3","and C4"],
					"filter_vals":["V1","V2"],
					"filter_ops":["=","!="]
				}
}`
	req1 := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(simpleBody))
	req1.Header.Set("Content-Type", "application/json")
	req2 := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(bodyWithFilter))
	req2.Header.Set("Content-Type", "application/json")
	req3 := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(bodyFailWithOpsValidError))
	req3.Header.Set("Content-Type", "application/json")
	req4 := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(bodyFailWithValidError))
	req4.Header.Set("Content-Type", "application/json")
	p1, err1 := services.CreateSqlParams(req1)
	p2, err2 := services.CreateSqlParams(req2)
	_, err3 := services.CreateSqlParams(req3)
	_, err4 := services.CreateSqlParams(req4)
	tr := true
	a1 := entities.SqlParams{
		SelectCols: []string{"Col1", "Col2"},
		Table:      "t1",
		Distinct:   &tr,
	}
	f := entities.FilterParams{
		FilterCols: []string{"C3", "and C4"},
		FilterVals: []string{"V1", "V2"},
		FilterOps:  []string{"=", "<="},
	}
	a2 := entities.SqlParams{
		SelectCols: []string{"Col1", "Col2"},
		Table:      "t1",
		Distinct:   &tr,
		Filters:    &f,
	}
	assert.Equal(t, err1, nil)
	assert.Equal(t, err2, nil)
	assert.Error(t, err3)
	assert.Error(t, err4)
	assert.Equal(t, p1, a1)
	assert.Equal(t, p2, a2)

}
func TestSqlStatementCreator(t *testing.T) {
	tr := true
	fl := false
	dialect_default := ""
	dialect_pg := "pg"
	dialect_sqlserver := "sqlserver"
	limit := "2"
	s1 := entities.SortParams{
		SortCols: []string{"s1", "s2"},
		SortType: []string{"asc", "desc"},
	}
	s2 := entities.SortParams{
		SortCols: []string{"s1"},
		SortType: []string{"asc"},
	}
	f := entities.FilterParams{
		FilterCols: []string{"C3", "and C4"},
		FilterVals: []string{"'V1'", "V2"},
		FilterOps:  []string{"=", "<="},
	}
	sql_params1 := entities.SqlParams{
		SelectCols: []string{"Col1", "Col2"},
		Table:      "t1",
		Distinct:   &tr,
		Filters:    &f,
	}
	sql_params2 := entities.SqlParams{
		SelectCols: []string{"Col1", "Col2"},
		Table:      "t1",
		Distinct:   &tr,
		Filters:    &f,
		LimitVals:  &limit,
	}
	sql_params3 := entities.SqlParams{
		SelectCols: []string{"Col1", "Col2"},
		Table:      "t1",
		Distinct:   &tr,
		Filters:    &f,
		LimitVals:  &limit,
		Sort:       &s1,
	}
	sql_params4 := entities.SqlParams{
		SelectCols: []string{"Col1", "Col2"},
		Table:      "t1",
		Distinct:   &tr,
		Filters:    &f,
		LimitVals:  &limit,
		Sort:       &s2,
	}
	sql_params5 := entities.SqlParams{
		SelectCols: []string{"Col1", "Col2"},
		Table:      "t1",
		Distinct:   &fl,
		Filters:    &f,
		LimitVals:  &limit,
		Sort:       &s2,
	}
	stmt_1 := services.GetStatementSql(&dialect_pg, sql_params1)
	stmt_2 := services.GetStatementSql(&dialect_pg, sql_params2)
	stmt_3 := services.GetStatementSql(&dialect_pg, sql_params3)
	stmt_4 := services.GetStatementSql(&dialect_pg, sql_params4)
	stmt_5 := services.GetStatementSql(&dialect_sqlserver, sql_params1) //Sql server test
	stmt_6 := services.GetStatementSql(&dialect_sqlserver, sql_params2)
	stmt_7 := services.GetStatementSql(&dialect_sqlserver, sql_params3)
	stmt_8 := services.GetStatementSql(&dialect_sqlserver, sql_params4)
	stmt_9 := services.GetStatementSql(&dialect_sqlserver, sql_params5)
	stmt_10 := services.GetStatementSql(&dialect_pg, sql_params5)
	stmt_11 := services.GetStatementSql(&dialect_default, sql_params5)

	actual_1 := `Select Distinct Col1,Col2 from t1 WHERE C3 = 'V1' and C4 <= V2`
	actual_2 := `Select Distinct Col1,Col2 from t1 WHERE C3 = 'V1' and C4 <= V2 Limit 2`
	actual_3 := `Select Distinct Col1,Col2 from t1 WHERE C3 = 'V1' and C4 <= V2 ORDER BY s1 asc, s2 desc Limit 2`
	actual_4 := `Select Distinct Col1,Col2 from t1 WHERE C3 = 'V1' and C4 <= V2 ORDER BY s1 asc Limit 2`
	actual_6 := `Select Distinct TOP 2 Col1,Col2 from t1 WHERE C3 = 'V1' and C4 <= V2`
	actual_7 := `Select Distinct TOP 2 Col1,Col2 from t1 WHERE C3 = 'V1' and C4 <= V2 ORDER BY s1 asc, s2 desc`
	actual_8 := `Select Distinct TOP 2 Col1,Col2 from t1 WHERE C3 = 'V1' and C4 <= V2 ORDER BY s1 asc`
	actual_9 := `Select TOP 2 Col1,Col2 from t1 WHERE C3 = 'V1' and C4 <= V2 ORDER BY s1 asc`
	actual_10 := `Select Col1,Col2 from t1 WHERE C3 = 'V1' and C4 <= V2 ORDER BY s1 asc Limit 2`

	assert.Equal(t, stmt_1, actual_1)
	assert.Equal(t, stmt_2, actual_2)
	assert.Equal(t, stmt_3, actual_3)
	assert.Equal(t, stmt_4, actual_4)
	assert.Equal(t, stmt_5, actual_1) //Sql server test case
	assert.Equal(t, stmt_6, actual_6)
	assert.Equal(t, stmt_7, actual_7)
	assert.Equal(t, stmt_8, actual_8)
	assert.Equal(t, stmt_9, actual_9)
	assert.Equal(t, stmt_10, actual_10)
	assert.Equal(t, stmt_11, actual_10)
}
