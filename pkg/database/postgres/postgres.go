package postgres

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/julioshinoda/challenge-api/pkg/database"
)

// DBConn is the struct that implements SQLInterface
type DBConn struct{}

// GetDBConn returns the struct that implements SQLInterface
func GetDBConn() database.SQLInterface {
	return DBConn{}
}

func createConnection() (*pgx.Conn, error) {
	conn, connErr := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	return conn, connErr
}

// QueryExecutor opens a connection, execute query that it receives and closes the connection
func (psql DBConn) QueryExecutor(config database.QueryConfig) ([]interface{}, error) {
	var resArray []interface{}

	conn, connErr := createConnection()
	if connErr != nil {
		return nil, connErr
	}
	defer conn.Close(context.Background())

	rows, queryErr := conn.Query(context.Background(), config.QueryStr, config.Values...)

	for rows.Next() {
		value, valueErr := rows.Values()
		if valueErr != nil {
			return []interface{}{}, valueErr
		}

		resArray = append(resArray, value)
	}

	return resArray, queryErr

}

// TransactionExecutor opens a connection, executes a transaction and closes the connection
func (psql DBConn) TransactionExecutor(configs []database.QueryConfig) error {
	conn, connErr := createConnection()
	if connErr != nil {
		return connErr
	}
	defer conn.Close(context.Background())

	tx, beginErr := conn.Begin(context.Background())
	if beginErr != nil {
		return beginErr
	}
	defer tx.Rollback(context.Background())
	// Rollback is safe to call even if the tx is already closed, so if the tx commits successfully, this is a no-op

	for _, config := range configs {
		_, execErr := tx.Exec(context.Background(), config.QueryStr, config.Values...)
		if execErr != nil {
			return execErr
		}
	}

	commitErr := tx.Commit(context.Background())
	if commitErr != nil {
		return commitErr
	}

	return nil
}
