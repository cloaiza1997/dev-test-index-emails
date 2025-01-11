package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	email "00-indexer/functions/emails"
	zs "00-indexer/functions/zincsearch"
)

type EmailError struct {
	Error string `json:"error"`
	Path  string `json:"path"`
}

func main() {
	startTime := time.Now()

	mailDir := "./mock/maildir"

	ok, successCount, errorCount := uploadEmails(mailDir)

	duration := time.Since(startTime)

	fmt.Printf("Ok: %t - Duration: %v => Success: %d | Error: %d\n", ok, duration, successCount, errorCount)
}

func uploadEmails(mailDir string) (bool, int, int) {
	emails := []email.Email{}
	emailErrors := []EmailError{}

	total := 0
	totalErrors := 0
	totalSuccess := 0

	emailsCh := make(chan struct{}, 10)
	var mtx sync.Mutex

	err := filepath.Walk(mailDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error walking the path (%s): %v", mailDir, err)
		}

		if !info.IsDir() {
			emailsCh <- struct{}{}

			go func() {
				defer func() { <-emailsCh }()

				emailJson, err := email.ParseEmailFile(path)

				mtx.Lock()

				total++
				hasError := err != nil

				if hasError {
					totalErrors++
					errorMessage := fmt.Sprintf("Error parsing email: %v", err)
					emailErrors = append(emailErrors, EmailError{Error: errorMessage, Path: path})
				} else {
					totalSuccess++
					emails = append(emails, emailJson)
				}

				fmt.Printf("Email %d - %t => %s\n", total, !hasError, path)

				mtx.Unlock()
			}()
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error proccessing emails => %v\n", err)

		return false, totalSuccess, totalErrors
	}

	for i := 0; i < cap(emailsCh); i++ {
		emailsCh <- struct{}{}
	}

	if len(emailErrors) > 0 {
		fmt.Printf("Errors parsing emails: %v\n", emailErrors)
	}

	ok := bulkEmails(emails)

	return ok, totalSuccess, totalErrors
}

func bulkEmails(emails []email.Email) bool {
	data := zs.Bulk[email.Email]{Index: "emails", Records: emails}

	ok, message := zs.BulkPost(data)

	fmt.Printf("Result uploading emails to ZincSearch (%t)\n", ok)
	fmt.Printf("%v\n", message)

	return ok
}
