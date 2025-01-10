package functions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var ZS_USER = os.Getenv("EMAIL_INDEX_ZS_USER")
var ZS_PASS = os.Getenv("EMAIL_INDEX_ZS_PASS")
var ZS_HOST = "http://localhost:4080/api"

type Bulk[T any] struct {
	Index   string `json:"index"`
	Records []T    `json:"records"`
}

func BulkPost[T any](data Bulk[T]) error {
	jsonData, err := json.Marshal(data)

	if err != nil {
		return fmt.Errorf("error marshalling data (index=%s): %v", data.Index, err)
	}

	request, err := http.NewRequest("POST", ZS_HOST+"/_bulkv2", bytes.NewReader(jsonData))

	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth(ZS_USER, ZS_PASS)

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("error response: %v", response)
	}

	fmt.Printf("OK - %d records uploaded: %v\n", len(data.Records), response)

	return nil
}
