package functions

import (
	"fmt"
	"os"
	"sync"
	"time"

	utils "github.com/cloaiza1997/dev-test-tr-emails/functions/utils"
)

var ZS_USER = os.Getenv("ZINC_FIRST_ADMIN_USER")
var ZS_PASS = os.Getenv("ZINC_FIRST_ADMIN_PASSWORD")
var ZS_HOST = os.Getenv("EMAIL_INDEX_ZS_HOST")

type Bulk[T any] struct {
	Index   string `json:"index"`
	Records []T    `json:"records"`
}

type PropertyDetail struct {
	Type          string `json:"type"`
	Index         bool   `json:"index"`
	Store         bool   `json:"store"`
	Sortable      bool   `json:"sortable"`
	Aggregatable  bool   `json:"aggregatable"`
	Highlightable bool   `json:"highlightable"`
}

type Mapping struct {
	Properties map[string]PropertyDetail `json:"properties"`
}

type IndexStructure struct {
	Name         string  `json:"name"`
	StorageType  string  `json:"storage_type"`
	ShardNum     int     `json:"shard_num"`
	MappingField Mapping `json:"mappings"`
}

func BulkRecords[T any](index string, records []T, zincSearchLogs *[]string) {
	data := Bulk[T]{Index: index, Records: records}

	ok, message := bulkBatch(data)

	log := fmt.Sprintf("Result uploading %s to ZincSearch (%t) => %v", index, ok, message)

	*zincSearchLogs = append(*zincSearchLogs, log)
}

func CreateIndex(data IndexStructure) (bool, error) {
	fmt.Println("ZincSearch create index...")

	response, err := utils.NewRequest(utils.Request{
		Method: "POST",
		Url:    ZS_HOST + "/index",
		Body:   data,
		Auth: utils.RequestAuth{
			User: ZS_USER,
			Pass: ZS_PASS,
		},
	})

	if err != nil {
		return false, err
	}

	return response.StatusCode == 200, nil
}

func IndexBatchZincSearch[T any](index string, batchEmails [][]T, zincSearchLogs *[]string, wg *sync.WaitGroup, ch chan struct{}) {
	startTime := time.Now()

	for _, emails := range batchEmails {
		wg.Add(1)
		ch <- struct{}{}

		go func() {
			defer func() {
				<-ch
				wg.Done()
			}()

			BulkRecords(index, emails, zincSearchLogs)
		}()
	}

	wg.Wait()

	duration := time.Since(startTime)

	fmt.Println("Duration uploading emails:", duration)
}

func ValidateIndexExists(index string) (bool, error) {
	fmt.Println("ZincSearch validate index...")

	response, err := utils.NewRequest(utils.Request{
		Method: "GET",
		Url:    ZS_HOST + "/" + index + "/_mapping",
		Auth: utils.RequestAuth{
			User: ZS_USER,
			Pass: ZS_PASS,
		},
	})

	if response.StatusCode != 400 && err != nil {
		return false, err
	}

	return response.StatusCode == 200, nil
}

func bulkBatch[T any](data Bulk[T]) (bool, string) {
	fmt.Println("ZincSearch bulk upload...")

	startTime := time.Now()

	response, err := utils.NewRequest(utils.Request{
		Method: "POST",
		Url:    ZS_HOST + "/_bulkv2",
		Body:   data,
		Auth: utils.RequestAuth{
			User: ZS_USER,
			Pass: ZS_PASS,
		},
	})

	if err != nil {
		return false, fmt.Sprintf("%v", err)
	}

	duration := time.Since(startTime)

	return true, fmt.Sprintf("ZincSearch bulk upload - %v - Payload %d => %s", duration, len(data.Records), response.Status)
}
