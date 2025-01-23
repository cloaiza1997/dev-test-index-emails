package functions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Request struct {
	Method string
	Url    string
	Body   interface{}
	Auth   RequestAuth
}

type RequestAuth struct {
	User string
	Pass string
}

func FormatTime() (time.Time, string) {
	now := time.Now()

	return now, now.Format("2006-01-02 15:04:05.000")
}

func GetJsonData[T any](filepath string) (T, error) {
	var indexerData T

	file, err := os.Open(filepath)

	if err != nil {
		return indexerData, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&indexerData)

	if err != nil {
		return indexerData, err
	}

	return indexerData, nil
}

func Log(message string) {
	_, timeFormated := FormatTime()

	fmt.Printf("%s - %s\n", timeFormated, message)
}

func NewRequest(options Request) (http.Response, error) {
	var emptyResponse http.Response
	jsonData, err := json.Marshal(options.Body)

	if err != nil {
		return emptyResponse, fmt.Errorf("error parsing body: %v", err)
	}

	request, err := http.NewRequest(options.Method, options.Url, bytes.NewReader(jsonData))

	if err != nil {
		return emptyResponse, fmt.Errorf("error creating request: %v", err)
	}

	request.Header.Set("Content-Type", "application/json")

	if options.Auth.User != "" && options.Auth.Pass != "" {
		request.SetBasicAuth(options.Auth.User, options.Auth.Pass)
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return emptyResponse, fmt.Errorf("error do the request: %v", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return *response, fmt.Errorf("error response: %v", response)
	}

	return *response, nil
}
