package functions

import (
	"fmt"
	"sync"
	"time"

	email "github.com/cloaiza1997/dev-test-tr-emails/functions/emails"
	fs "github.com/cloaiza1997/dev-test-tr-emails/functions/files"
	util "github.com/cloaiza1997/dev-test-tr-emails/functions/utils"
	zs "github.com/cloaiza1997/dev-test-tr-emails/functions/zincsearch"
)

const INDEX = "emails"
const BATCH_SIZE = 10000

func InitUpload(mailDir string, indexByBatch bool) {
	startTime, startTimeFormated := util.FormatTime()

	fmt.Printf("%s - Start indexing emails (Emails=%s, Batch=%t)...\n", startTimeFormated, mailDir, indexByBatch)

	ok, successCount, errorCount, logs := uploadEmails(mailDir, indexByBatch)

	for i, log := range logs {
		fmt.Println(i, log)
	}

	duration := time.Since(startTime)

	_, endTimeFormated := util.FormatTime()

	fmt.Printf("%s - Duration: %v => Ok: %t | Parsed Success: %d | Parsed Error: %d\n", endTimeFormated, duration, ok, successCount, errorCount)
}

func uploadEmails(mailDir string, indexByBatch bool) (bool, int, int, []string) {
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

	errCount := fs.WalkFilePath(mailDir, func(path string) { totalEmails++ })

	if errCount != nil {
		return handleReturnError(getErrorMessage(errCount))
	}

	fmt.Printf("Total emails: %d\n", totalEmails)

	errMails := fs.WalkFilePath(mailDir, func(path string) {
		email.HandleFile(email.HandleFileOptions{
			Path:                path,
			IndexByBatch:        indexByBatch,
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

	if errMails != nil {
		return handleReturnError(getErrorMessage(errMails))
	}

	if len(emailErrors) > 0 {
		fmt.Println(getErrorMessage(emailErrors))
	}

	zs.IndexBatchZincSearch(INDEX, batchEmails, &zincSearchLogs, &wg, emailsCh)

	totalEmailErrors := len(emailErrors)
	totalEmailSuccess := totalEmails - totalEmailErrors

	return true, totalEmailSuccess, totalEmailErrors, zincSearchLogs
}

func getErrorMessage(err any) string {
	return fmt.Sprintf("Error proccessing emails => %v", err)
}

func handleReturnError(message string) (bool, int, int, []string) {
	fmt.Println(message)

	return false, 0, 0, []string{}
}
