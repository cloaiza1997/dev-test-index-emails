package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type Email struct {
	MessageID               string `json:"messageId"`
	Date                    string `json:"date"`
	From                    string `json:"from"`
	To                      string `json:"to"`
	Subject                 string `json:"subject"`
	MimeVersion             string `json:"mimeVersion"`
	ContentType             string `json:"contentType"`
	ContentTransferEncoding string `json:"contentTransferEncoding"`
	XFrom                   string `json:"xFrom"`
	XTo                     string `json:"xTo"`
	XCc                     string `json:"xCc"`
	XBcc                    string `json:"xBcc"`
	XFolder                 string `json:"xFolder"`
	XOrigin                 string `json:"xOrigin"`
	XFileName               string `json:"xFileName"`
	Body                    string `json:"body"`
	Path                    string `json:"path"`
}

func main() {
	startTime := time.Now()

	mailDir := "./maildir"

	fmt.Println("Hello, World!")

	emails, _ := GetMails(mailDir)

	uploadEmails(emails)

	duration := time.Since(startTime)

	fmt.Println("Duration", duration)
}

func GetMails(mailDir string) ([]Email, error) {
	var emails []Email
	var mtx sync.Mutex

	emailsCh := make(chan struct{}, 200)

	err := filepath.Walk(mailDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			emailsCh <- struct{}{}

			go func() {
				defer func() { <-emailsCh }()

				email, _ := parseEmailFile(path)

				mtx.Lock()
				emails = append(emails, email)
				mtx.Unlock()
			}()
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error walking the path")
		return nil, err
	}

	for i := 0; i < cap(emailsCh); i++ {
		emailsCh <- struct{}{}
	}

	return emails, nil
}

func parseEmailFile(path string) (Email, error) {
	fileContent, err := os.ReadFile(path)

	if err != nil {
		return Email{}, err
	}

	lines := strings.Split(string(fileContent), "\n")

	var email Email

	mapper := make(map[string]*string)

	mapper["Message-ID"] = &email.MessageID
	mapper["Date"] = &email.Date
	mapper["From"] = &email.From
	mapper["To"] = &email.To
	mapper["Subject"] = &email.Subject
	mapper["Mime-Version"] = &email.MimeVersion
	mapper["Content-Type"] = &email.ContentType
	mapper["Content-Transfer-Encoding"] = &email.ContentTransferEncoding
	mapper["X-From"] = &email.XFrom
	mapper["X-To"] = &email.XTo
	mapper["X-cc"] = &email.XCc
	mapper["X-bcc"] = &email.XBcc
	mapper["X-Folder"] = &email.XFolder
	mapper["X-Origin"] = &email.XOrigin
	mapper["X-FileName"] = &email.XFileName

	bodyLines := []string{}

	email.Path = path

	for _, line := range lines {
		prefix := strings.Split(line, ":")[0]

		value, ok := mapper[prefix]

		if ok && strings.TrimSpace(*value) == "" {
			*value = strings.TrimSpace(strings.TrimPrefix(line, prefix+":"))
		} else {
			bodyLines = append(bodyLines, line)
		}
	}

	email.Body = strings.TrimSpace(strings.Join(bodyLines, "\n"))

	return email, nil
}

func uploadEmails(emails []Email) error {

	type EmailIndex struct {
		Index   string  `json:"index"`
		Records []Email `json:"records"`
	}

	data := EmailIndex{Index: "emails", Records: emails}

	jsonData, _ := json.Marshal(data)

	request, err := http.NewRequest("POST", "http://localhost:4080/api/_bulkv2", bytes.NewReader(jsonData))

	if err != nil {
		fmt.Println("Error creating request", err)
	}

	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth("admin", "Complexpass#123")

	client := &http.Client{}
	response, err2 := client.Do(request)

	if err2 != nil {
		fmt.Println("Error sending request", err2)
	} else {
		fmt.Printf("OK %d emails cargados - %v", len(emails), response)
	}

	return nil
}
