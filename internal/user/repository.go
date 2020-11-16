package user

import (
	"errors"

	"github.com/julioshinoda/challenge-api/pkg/database"
	"github.com/julioshinoda/challenge-api/pkg/database/postgres"
)

type Repository interface {
	GetUser(username string) (User, error)
}

type repo struct {
	DB database.SQLInterface
}

func NewRepo() Repository {
	return repo{DB: postgres.GetDBConn()}
}

func (r repo) GetUser(username string) (User, error) {
	query := "select id,username,secret from users where username = $1"
	result, err := r.DB.QueryExecutor(database.QueryConfig{QueryStr: query,
		Values: []interface{}{username}})

	if err != nil {
		return User{}, err
	}
	if len(result) == 1 {
		for _, row := range result {
			return User{
				ID:       row.([]interface{})[0].(int64),
				Username: row.([]interface{})[1].(string),
				Secret:   row.([]interface{})[2].(string),
			}, nil

		}
	}
	return User{}, errors.New("duplicated registry")
}
