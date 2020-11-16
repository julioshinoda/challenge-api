package expression

import (
	"errors"
	"reflect"
	"testing"

	"github.com/julioshinoda/challenge-api/internal/expression/mocks"
	"github.com/julioshinoda/challenge-api/internal/expression/model"
	"github.com/stretchr/testify/mock"
)

func Test_parseExpression(t *testing.T) {
	type args struct {
		expr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "[T0/parseExpression] success to parse expression",
			args: args{expr: "((x OR y) AND (z OR k) OR j)"},
			want: "((x || y) && (z || k) || j)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseExpression(tt.args.expr); got != tt.want {
				t.Errorf("parseExpression() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_AddExpression(t *testing.T) {
	type fields struct {
		Repo Repository
	}
	type args struct {
		expr model.Expression
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Expression
		wantErr bool
	}{
		{
			name: "[T0/AddExpression] save new expression",
			fields: fields{
				Repo: &mocks.Repository{},
			},
			args: args{
				expr: model.Expression{Expression: "x AND y"},
			},
			want: model.Expression{Expression: "x AND y", ID: int64(1)},
		},
		{
			name: "[T1/AddExpression] update new expression",
			fields: fields{
				Repo: &mocks.Repository{},
			},
			args: args{
				expr: model.Expression{Expression: "x OR y", ID: int64(1)},
			},
			want: model.Expression{Expression: "x OR y", ID: int64(1)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{
				Repo: tt.fields.Repo,
			}
			switch tt.name {
			case "[T0/AddExpression] save new expression":
				s.Repo.(*mocks.Repository).On("Create", mock.Anything).Return(model.Expression{Expression: "x AND y", ID: int64(1)}, nil)
			case "[T1/AddExpression] update new expression":
				s.Repo.(*mocks.Repository).On("Update", mock.Anything).Return(model.Expression{Expression: "x OR y", ID: int64(1)}, nil)
			}
			got, err := s.AddExpression(tt.args.expr)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.AddExpression() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.AddExpression() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Evaluate(t *testing.T) {
	type fields struct {
		Repo Repository
	}
	type args struct {
		expressionID int64
		params       map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "[T0/Evaluate] succes to evaluate",
			fields: fields{
				Repo: &mocks.Repository{},
			},
			args: args{
				expressionID: int64(1),
				params:       map[string]interface{}{"x": true, "y": true},
			},
			want: true,
		},
		{
			name: "[T1/Evaluate] error to get expression",
			fields: fields{
				Repo: &mocks.Repository{},
			},
			args: args{
				expressionID: int64(1),
				params:       map[string]interface{}{"x": true, "y": true},
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "[T2/Evaluate] error to evaluate expression",
			fields: fields{
				Repo: &mocks.Repository{},
			},
			args: args{
				expressionID: int64(1),
				params:       map[string]interface{}{"x": true, "y": true},
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{
				Repo: tt.fields.Repo,
			}
			switch tt.name {
			case "[T0/Evaluate] succes to evaluate":
				s.Repo.(*mocks.Repository).On("GetExpressionByID", mock.Anything).Return(model.Expression{Expression: "x AND y", ID: int64(1)}, nil)
			case "[T1/Evaluate] error to get expression":
				s.Repo.(*mocks.Repository).On("GetExpressionByID", mock.Anything).Return(model.Expression{}, errors.New("error on get expression"))
			case "[T2/Evaluate] error to evaluate expression":
				s.Repo.(*mocks.Repository).On("GetExpressionByID", mock.Anything).Return(model.Expression{Expression: "x sa y", ID: int64(1)}, nil)
			}
			got, err := s.Evaluate(tt.args.expressionID, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Evaluate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("service.Evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
