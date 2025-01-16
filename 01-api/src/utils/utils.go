package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/cloaiza1997/dev-test-tr-emails/src/models"
)

func DoResponse(w http.ResponseWriter, status int, data interface{}, err error) {
	response := models.Response{
		Success: status == http.StatusOK,
		Data:    data,
	}

	if err != nil {
		status = http.StatusInternalServerError
		response.Success = false
		response.Message = err.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(response)
}

func GetPagination(total, count, limit, page int) models.Pagination {
	if total == 0 {
		return models.Pagination{}
	}

	pages := int(math.Ceil(float64(total) / float64(limit)))

	prev := page - 1
	var next int

	if page < pages {
		next = page + 1
	} else {
		next = 0
	}

	return models.Pagination{
		Total: total,
		Count: count,
		Pages: pages,
		Prev:  prev,
		Next:  next,
	}
}

func GetQueryParam(r *http.Request, name string) string {
	return r.URL.Query().Get(name)
}

func GetQueryParamInt(r *http.Request, name string) int {
	value := GetQueryParam(r, name)

	valueNumber, _ := strconv.Atoi(value)

	return valueNumber
}

func NewRequest[T any](options models.Request) (T, error) {
	var result T
	jsonData, err := json.Marshal(options.Body)

	if err != nil {
		return logError[T]("Error parsing body", err)
	}

	request, err := http.NewRequest(options.Method, options.Url, bytes.NewReader(jsonData))

	if err != nil {
		return logError[T]("Error creating request", err)
	}

	request.Header.Set("Content-Type", "application/json")

	if options.Auth.User != "" && options.Auth.Pass != "" {
		request.SetBasicAuth(options.Auth.User, options.Auth.Pass)
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return logError[T]("Error do the request", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return logError[T]("Error response", response)
	}

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return logError[T]("Error decoding response", err)
	}

	return result, nil
}

func logError[T any](message string, err interface{}) (T, error) {
	var result T

	newError := fmt.Errorf("%s: %v", message, err)

	fmt.Println(newError)

	return result, newError
}
