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

func main() {
	startTime := time.Now()

	mailDir := "./mock/maildir"

	result := uploadEmails(mailDir)

	fmt.Println(result)

	duration := time.Since(startTime)

	fmt.Printf("Duration: %v\n", duration)
}

func uploadEmails(mailDir string) bool {
	var emails []email.Email
	var mtx sync.Mutex

	emailsCh := make(chan struct{}, 10)

	type EmailError struct {
		Error string `json:"error"`
		Path  string `json:"path"`
	}

	emailErrors := []EmailError{}
	i := 1

	err := filepath.Walk(mailDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error walking the path (%s): %v", mailDir, err)
		}

		if !info.IsDir() {
			emailsCh <- struct{}{}

			go func() {
				defer func() { <-emailsCh }()

				email, err := email.ParseEmailFile(path)

				mtx.Lock()

				log := fmt.Sprintf("Processing email %d: %s = ", i, path)

				if err != nil {
					log += "ERROR"

					errorMessage := fmt.Sprintf("Error parsing email: %v", err)
					emailErrors = append(emailErrors, EmailError{Error: errorMessage, Path: path})
				} else {
					log += "OK"

					emails = append(emails, email)
				}

				fmt.Println(log)
				i++

				mtx.Unlock()
			}()
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error proccessing emails: v%v\n", err)

		return false
	}

	for i := 0; i < cap(emailsCh); i++ {
		emailsCh <- struct{}{}
	}

	data := zs.Bulk[email.Email]{Index: "emails", Records: emails}

	zsErr := zs.BulkPost(data)

	if zsErr != nil {
		fmt.Printf("Error uploading emails to ZincSearch: %v\n", zsErr)

		return false
	}

	return true
}
