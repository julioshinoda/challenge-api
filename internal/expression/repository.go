package expression

import (
	"errors"
	"time"

	"github.com/julioshinoda/challenge-api/internal/expression/model"
	"github.com/julioshinoda/challenge-api/pkg/database"
	"github.com/julioshinoda/challenge-api/pkg/database/postgres"
)

type Repository interface {
	Create(expr model.Expression) (model.Expression, error)
	Update(expr model.Expression) (model.Expression, error)
	GetExpressions() ([]model.Expression, error)
	GetExpressionByID(ID int64) (model.Expression, error)
	DeleteExpressionByID(ID int64) error
}

type Repo struct {
	DB database.SQLInterface
}

func NewRepo() Repository {
	return Repo{DB: postgres.GetDBConn()}
}

func (r Repo) Create(expr model.Expression) (model.Expression, error) {
	query := `INSERT INTO expressions ("expression",created_at) VALUES ($1,$2) RETURNING id;`
	result, err := r.DB.QueryExecutor(database.QueryConfig{QueryStr: query,
		Values: []interface{}{expr.Expression, time.Now().Format(time.RFC3339)}})

	if err != nil {
		return expr, err
	}
	if len(result) == 0 {
		return expr, errors.New("cannot save this expression")
	}
	expr.ID = result[0].([]interface{})[0].(int64)
	return expr, nil
}

func (r Repo) Update(expr model.Expression) (model.Expression, error) {
	query := `UPDATE  expressions SET expression = $1 WHERE id = $2 RETURNING id;`
	result, err := r.DB.QueryExecutor(database.QueryConfig{QueryStr: query,
		Values: []interface{}{expr.Expression, expr.ID}})

	if err != nil {
		return expr, err
	}
	if len(result) == 0 {
		return expr, errors.New("cannot update this expression")
	}
	expr.ID = result[0].([]interface{})[0].(int64)
	return expr, nil
}

func (r Repo) GetExpressions() ([]model.Expression, error) {
	query := `SELECT id,expression FROM expressions`
	result, err := r.DB.QueryExecutor(database.QueryConfig{QueryStr: query,
		Values: []interface{}{}})
	expressionList := []model.Expression{}
	if err != nil {
		return expressionList, err
	}
	if len(result) == 0 {
		return expressionList, errors.New("not find any expression")
	}
	if len(result) > 0 {
		for _, row := range result {
			expression := model.Expression{
				ID:         row.([]interface{})[0].(int64),
				Expression: row.([]interface{})[1].(string),
			}
			expressionList = append(expressionList, expression)
		}
		return expressionList, nil
	}
	return expressionList, nil
}

func (r Repo) GetExpressionByID(ID int64) (model.Expression, error) {
	query := `SELECT expression FROM expressions WHERE id = $1`
	result, err := r.DB.QueryExecutor(database.QueryConfig{QueryStr: query,
		Values: []interface{}{ID}})
	if err != nil {
		return model.Expression{}, err
	}
	if len(result) == 1 {
		for _, row := range result {
			return model.Expression{
				ID:         ID,
				Expression: row.([]interface{})[0].(string),
			}, nil
		}
	}
	return model.Expression{}, errors.New("not found expression")
}

func (r Repo) DeleteExpressionByID(ID int64) error {
	query := `DELETE FROM expressions WHERE id = $1`
	_, err := r.DB.QueryExecutor(database.QueryConfig{QueryStr: query,
		Values: []interface{}{ID}})
	if err != nil {
		return err
	}
	return nil
}
