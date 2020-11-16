package expression

import (
	"strings"

	"github.com/PaesslerAG/gval"
	"github.com/julioshinoda/challenge-api/internal/expression/model"
)

type Manager interface {
	AddExpression(expr model.Expression) (model.Expression, error)
	GetAll() ([]model.Expression, error)
	Evaluate(expressionID int64, params map[string]interface{}) (bool, error)
	DeleteExpression(ID int64) error
}

type service struct {
	Repo Repository
}

func NewExpression() Manager {
	return service{Repo: NewRepo()}
}

func (s service) AddExpression(expr model.Expression) (model.Expression, error) {
	if expr.ID != int64(0) {
		return s.Repo.Update(expr)
	}
	return s.Repo.Create(expr)
}

func (s service) GetAll() ([]model.Expression, error) {
	return s.Repo.GetExpressions()
}

func (s service) Evaluate(expressionID int64, params map[string]interface{}) (bool, error) {
	expr, err := s.Repo.GetExpressionByID(expressionID)
	if err != nil {
		return false, err
	}
	value, err := gval.Evaluate(parseExpression(expr.Expression), params)
	if err != nil {
		return false, err
	}
	return value.(bool), nil
}

func (s service) DeleteExpression(ID int64) error {
	return s.Repo.DeleteExpressionByID(ID)
}

func parseExpression(expr string) string {
	expr = strings.ReplaceAll(expr, "OR", "||")
	expr = strings.ReplaceAll(expr, "AND", "&&")
	return expr
}
