package functions

import (
	"fmt"
	"sync"
	"time"

	email "github.com/cloaiza1997/dev-test-tr-emails/functions/emails"
	fs "github.com/cloaiza1997/dev-test-tr-emails/functions/files"
	utils "github.com/cloaiza1997/dev-test-tr-emails/functions/utils"
	zs "github.com/cloaiza1997/dev-test-tr-emails/functions/zincsearch"
)

const INDEX_STRUCTURE = "./data/index-structure.json"

type UploadOptions struct {
	BatchSize    int
	Index        string
	IndexByBatch bool
	MailDir      string
	Routines     int
}

func InitUpload(options UploadOptions) {
	startTime, startTimeFormated := utils.FormatTime()

	fmt.Printf("%s - Start indexing emails (Emails=%s, Batch=%t)...\n", startTimeFormated, options.MailDir, options.IndexByBatch)

	ok, successCount, errorCount, logs := uploadEmails(options)

	for i, log := range logs {
		utils.Log(fmt.Sprintf("[%d] %s", i, log))
	}

	duration := time.Since(startTime)

	_, endTimeFormated := utils.FormatTime()

	fmt.Printf("%s - Duration: %v => Ok: %t | Parsed Success: %d | Parsed Error: %d\n", endTimeFormated, duration, ok, successCount, errorCount)
}

func getErrorMessage(err any) string {
	return fmt.Sprintf("Error proccessing emails => %v", err)
}

func handleIndex(index string) error {
	exists, err := zs.ValidateIndexExists(index)

	if err != nil {
		return err
	}

	if !exists {
		data, err := utils.GetJsonData[zs.IndexStructure](INDEX_STRUCTURE)

		if err != nil {
			return err
		}

		data.Name = index

		ok, err := zs.CreateIndex(data)

		if err != nil {
			return err
		}

		if !ok {
			return fmt.Errorf("error creating index")
		}
	}

	return nil
}

func handleReturnError(message string) (bool, int, int, []string) {
	utils.Log(message)

	return false, 0, 0, []string{}
}

func uploadEmails(options UploadOptions) (bool, int, int, []string) {
	err := handleIndex(options.Index)

	if err != nil {
		return handleReturnError(getErrorMessage(err))
	}

	utils.Log("Count files ...")

	batchEmails := [][]email.Email{}
	emails := []email.Email{}
	emailErrors := []email.EmailError{}
	zincSearchLogs := []string{}

	totalEmails := 0
	totalBatch := 1
	totalEmailBatch := 0
	totalEmailProcessed := 0

	emailsCh := make(chan struct{}, options.Routines)
	var mtx sync.Mutex
	var wg sync.WaitGroup

	errCount := fs.WalkFilePath(options.MailDir, func(path string) { totalEmails++ })

	if errCount != nil {
		return handleReturnError(getErrorMessage(errCount))
	}

	utils.Log((fmt.Sprintf("Total emails: %d", totalEmails)))
	utils.Log("Processing files ...")

	errMails := fs.WalkFilePath(options.MailDir, func(path string) {
		email.HandleFile(email.HandleFileOptions{
			Path:                path,
			IndexByBatch:        options.IndexByBatch,
			BatchSize:           options.BatchSize,
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
		utils.Log(getErrorMessage(emailErrors))
	}

	zs.IndexBatchZincSearch(options.Index, batchEmails, &zincSearchLogs, &wg, emailsCh)

	totalEmailErrors := len(emailErrors)
	totalEmailSuccess := totalEmails - totalEmailErrors

	return true, totalEmailSuccess, totalEmailErrors, zincSearchLogs
}
