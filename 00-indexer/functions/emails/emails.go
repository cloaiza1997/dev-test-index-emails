package functions

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

type Email struct {
	MessageID               string `json:"messageId"`
	Date                    string `json:"date"`
	From                    string `json:"from"`
	To                      string `json:"to"`
	Cc                      string `json:"cc"`
	Bcc                     string `json:"bcc"`
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

type EmailError struct {
	Error string `json:"error"`
	Path  string `json:"path"`
}

type HandleFileOptions struct {
	IndexByBatch        bool
	BatchSize           int
	Path                string
	Ch                  *chan struct{}
	Wg                  *sync.WaitGroup
	Mtx                 *sync.Mutex
	BatchEmails         *[][]Email
	EmailErrors         *[]EmailError
	Emails              *[]Email
	TotalBatch          *int
	TotalEmailBatch     *int
	TotalEmailProcessed *int
	TotalEmails         *int
}

func HandleFile(options HandleFileOptions) {
	*options.Ch <- struct{}{}
	options.Wg.Add(1)

	go func() {
		defer func() {
			<-*options.Ch
			options.Wg.Done()
		}()

		emailJson, err := ParseEmailFile(options.Path)

		options.Mtx.Lock()

		hasError := err != nil

		if hasError {
			errorMessage := fmt.Sprintf("Error parsing email: %v", err)
			*options.EmailErrors = append(*options.EmailErrors, EmailError{Error: errorMessage, Path: options.Path})
		} else {
			*options.Emails = append(*options.Emails, emailJson)
		}

		*options.TotalEmailProcessed++
		*options.TotalEmailBatch++

		isLast := *options.TotalEmailProcessed == *options.TotalEmails

		if *options.TotalEmailProcessed%options.BatchSize == 0 || isLast {
			fmt.Printf("Batch %d (%d, %d) ...\n", *options.TotalBatch, *options.TotalEmailProcessed, *options.TotalEmailBatch)

			*options.TotalBatch++
			*options.TotalEmailBatch = 0

			if options.IndexByBatch || isLast {
				*options.BatchEmails = append(*options.BatchEmails, *options.Emails)
				*options.Emails = []Email{}
			}
		}

		options.Mtx.Unlock()
	}()
}

func ParseEmailFile(path string) (Email, error) {
	fileContent, err := os.ReadFile(path)

	if err != nil {
		return Email{}, fmt.Errorf("error reading file (%s): %v", path, err)
	}

	fileLines := strings.Split(string(fileContent), "\n")
	mapper := make(map[string]*string)

	email := Email{Path: path}
	bodyLines := []string{}

	mapper["Message-ID"] = &email.MessageID
	mapper["Date"] = &email.Date
	mapper["From"] = &email.From
	mapper["To"] = &email.To
	mapper["Cc"] = &email.Cc
	mapper["Bcc"] = &email.Bcc
	mapper["To"] = &email.Bcc
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

	for _, line := range fileLines {
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
