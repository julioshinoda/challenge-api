package database

type QueryConfig struct {
	QueryStr string
	Values   []interface{}
}

type SQLInterface interface {
	QueryExecutor(config QueryConfig) ([]interface{}, error)
	TransactionExecutor(configs []QueryConfig) error
}
