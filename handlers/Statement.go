package handlers

import (
	"dataserver/entities"
	"dataserver/services"
	"encoding/json"
	"net/http"
)

type StatementHandler struct {
	Cnx          *entities.Cnx
	AuthProvider *services.ProducerTokenService
}

func (s *StatementHandler) ServeHttp(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Bad Request: only POST is allowed", http.StatusBadRequest)
		return
	}
	sql_params, err := services.CreateSqlParams(r)
	if err != nil {
		http.Error(w, "Bad Request: Couldn't create sql_param object", http.StatusBadRequest)
		return
	}
	sql_stmt := services.GetStatementSql(&s.Cnx.Type, sql_params)
	resp, err := services.GetDataFromQuery(s.Cnx, sql_stmt)
	if err != nil {
		http.Error(w, "Couldn't execute query", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
