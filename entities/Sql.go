package entities

type Operators struct {
	Eq  string `default0:"="`
	Gte string `default0:">="`
	Lte string `default0:"<="`
	Gt  string `default0:">"`
	Lt  string `default0:"<"`
	Neq string `default0:"<>"`
	In  string `default0:"IN"`
	Or  string `default0:"OR"`
	And string `default0:"AND"`
	Not string `default0:"NOT"`
}
type Clauses struct {
	Where    string `default0:"WHERE"`
	Select   string `default0:"SELECT"`
	Order    string `default0:"ORDER"`
	By       string `default0:"BY"`
	Distinct string `default0:"DISTINCT"`
	Limit    string `default0:"LIMIT"`
}

type FilterParams struct {
	FilterCols []string `json:"filter_cols"`
	FilterVals []string `json:"filter_vals"`
	FilterOps  []string `json:"filter_ops"`
}
type SortParams struct {
	SortCols []string `json:"sort_cols"`
	SortType []string `json:"sort_type"`
}
type SqlParams struct {
	SelectCols []string      `json:"select_cols" validate:"min=1"`
	Table      string        `json:"table" validate:"required"`
	LimitVals  *string       `json:"limit_vals,omitempty"`
	Filters    *FilterParams `json:"filter_params,omitempty"`
	Sort       *SortParams   `json:"sort_params,omitempty"`
	Distinct   *bool         `json:"distinct"`
}
