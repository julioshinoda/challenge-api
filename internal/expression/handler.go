package expression

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/julioshinoda/challenge-api/internal/expression/model"
)

func Handler(r chi.Router) {
	r.Post("/expressions", Create)
	r.Get("/expressions", ListAll)
	r.Get("/evaluate/{expression_id}", EvaluateByID)
	r.Delete("/expressions/{expression_id}", Delete)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var request model.Expression
	json.NewDecoder(r.Body).Decode(&request)
	service := NewExpression()
	expr, err := service.AddExpression(request)
	if err != nil {
		response, _ := json.Marshal(map[string]string{"error": err.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write(response)
		return
	}
	response, _ := json.Marshal(expr)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func ListAll(w http.ResponseWriter, r *http.Request) {
	service := NewExpression()
	expr, err := service.GetAll()
	if err != nil {
		response, _ := json.Marshal(map[string]string{"error": err.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write(response)
		return
	}
	response, _ := json.Marshal(expr)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func EvaluateByID(w http.ResponseWriter, r *http.Request) {
	requestID := chi.URLParam(r, "expression_id")
	expressionID, _ := strconv.Atoi(requestID)
	service := NewExpression()
	expr, err := service.Evaluate(int64(expressionID), parseEvaluateParams(r.URL.Query()))
	if err != nil {
		response, _ := json.Marshal(map[string]string{"error": err.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write(response)
		return
	}
	response, _ := json.Marshal(map[string]interface{}{"result": expr})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	requestID := chi.URLParam(r, "expression_id")
	expressionID, _ := strconv.Atoi(requestID)
	service := NewExpression()
	err := service.DeleteExpression(int64(expressionID))
	if err != nil {
		response, _ := json.Marshal(map[string]string{"error": err.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write(response)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func parseEvaluateParams(val url.Values) map[string]interface{} {
	params := map[string]interface{}{}
	for name, value := range val {
		result, _ := strconv.ParseBool(value[0])
		params[name] = result
	}
	return params
}
