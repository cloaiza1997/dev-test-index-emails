package main

import (
	"fmt"
	"sync"
	"time"

	email "github.com/cloaiza1997/dev-test-tr-emails/functions/emails"
	fs "github.com/cloaiza1997/dev-test-tr-emails/functions/files"
	zs "github.com/cloaiza1997/dev-test-tr-emails/functions/zincsearch"
)

var INDEX = "emails"
var INDEXING_BY_BATCH = true
var BATCH_SIZE = 10000

func main() {
	startTime, startTimeFormated := formatTime()

	fmt.Printf("%s - Start indexing emails...\n", startTimeFormated)

	mailDir := "./mock/maildir"

	ok, successCount, errorCount, logs := uploadEmails(mailDir)

	for i, log := range logs {
		fmt.Println(i, log)
	}

	duration := time.Since(startTime)

	_, endTimeFormated := formatTime()

	fmt.Printf("%s - Duration: %v => Ok: %t | Success: %d | Error: %d\n", endTimeFormated, duration, ok, successCount, errorCount)
}

func formatTime() (time.Time, string) {
	now := time.Now()

	return now, now.Format("2006-01-02 15:04:05")
}

func uploadEmails(mailDir string) (bool, int, int, []string) {
	batchEmails := [][]email.Email{}
	emails := []email.Email{}
	emailErrors := []email.EmailError{}
	zincSearchLogs := []string{}

	totalEmails := 0
	totalBatch := 1
	totalEmailBatch := 0
	totalEmailProcessed := 0

	emailsCh := make(chan struct{}, 10)
	var mtx sync.Mutex
	var wg sync.WaitGroup

	err := fs.WalkFilePath(mailDir, func(path string) { totalEmails++ })

	if err != nil {
		return handleError(fmt.Sprintf("Error proccessing emails => %v\n", err))
	}

	fmt.Printf("Total emails: %d\n", totalEmails)

	err = fs.WalkFilePath(mailDir, func(path string) {
		email.HandleFile(email.HandleFileOptions{
			Path:                path,
			IndexByBatch:        INDEXING_BY_BATCH,
			BatchSize:           BATCH_SIZE,
			Ch:                  &emailsCh,
			Wg:                  &wg,
			Mtx:                 &mtx,
			BatchEmails:         &batchEmails,
			EmailErrors:         &emailErrors,
			Emails:              &emails,
			TotalBatch:          &totalBatch,
			TotalEmailBatch:     &totalEmailBatch,
			TotalEmailProcessed: &totalEmailProcessed,
			TotalEmails:         &totalEmails,
		})
	})

	wg.Wait()

	if err != nil {
		return handleError(fmt.Sprintf("Error proccessing emails => %v\n", err))
	}

	zs.IndexBatchZincSearch(INDEX, batchEmails, &zincSearchLogs, &wg, emailsCh)

	totalEmailErrors := len(emailErrors)
	totalEmailSuccess := totalEmails - totalEmailErrors

	return true, totalEmailSuccess, totalEmailErrors, zincSearchLogs
}

func handleError(message string) (bool, int, int, []string) {
	fmt.Println(message)

	return false, 0, 0, []string{}
}
