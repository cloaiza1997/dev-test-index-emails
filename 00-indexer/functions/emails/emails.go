package functions

import (
	"bytes"
	"fmt"
	"io"
	"net/mail"
	"os"
	"strings"
	"sync"

	utils "github.com/cloaiza1997/dev-test-index-emails/functions/utils"
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
			utils.Log(fmt.Sprintf("Batch %d (%d, %d) ...", *options.TotalBatch, *options.TotalEmailProcessed, *options.TotalEmailBatch))

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
		return handleError(path, err)
	}

	email := Email{Path: path}

	return GetEmail(fileContent, &email)
}

func GetEmailByReader(fileContent []byte, email *Email) (Email, error) {
	reader := bytes.NewReader(fileContent)
	message, err := mail.ReadMessage(reader)

	if err != nil {
		return handleError(email.Path, err)
	}

	email.MessageID = getHeader("Message-ID", message)
	email.Date = getHeader("Date", message)
	email.From = getHeader("From", message)
	email.To = getHeader("To", message)
	email.Cc = getHeader("Cc", message)
	email.Bcc = getHeader("Bcc", message)
	email.Subject = getHeader("Subject", message)
	email.MimeVersion = getHeader("Mime-Version", message)
	email.ContentType = getHeader("Content-Type", message)
	email.ContentTransferEncoding = getHeader("Content-Transfer-Encoding", message)
	email.XFrom = getHeader("X-From", message)
	email.XTo = getHeader("X-To", message)
	email.XCc = getHeader("X-cc", message)
	email.XBcc = getHeader("X-bcc", message)
	email.XFolder = getHeader("X-Folder", message)
	email.XOrigin = getHeader("X-Origin", message)
	email.XFileName = getHeader("X-FileName", message)

	body, e := io.ReadAll(message.Body)

	if e != nil {
		return handleError(email.Path, err)
	}

	email.Body = strings.TrimSpace(string(body))

	return *email, nil
}

func GetEmail(fileContent []byte, email *Email) (Email, error) {
	fileLines := strings.Split(string(fileContent), "\n")
	mapper := make(map[string]*string)

	bodyLines := []string{}

	mapper["Message-ID"] = &email.MessageID
	mapper["Date"] = &email.Date
	mapper["From"] = &email.From
	mapper["To"] = &email.To
	mapper["Cc"] = &email.Cc
	mapper["Bcc"] = &email.Bcc
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

	for i := 0; i < len(fileLines); i++ {
		line := fileLines[i]

		if len(line) == 0 || line[0:1] == "\r" {
			bodyLines = append(bodyLines, fileLines[i:]...)
			break
		}

		header := strings.Split(line, ":")

		if isEmpty(line[0:1]) || len(header) <= 1 {
			return Email{}, fmt.Errorf("invalid line: %q", line)
		}

		setHeaderValue(&i, &line, &fileLines)

		prefix := header[0]

		value, ok := mapper[prefix]

		if ok && strings.TrimSpace(*value) == "" {
			*value = strings.TrimSpace(strings.TrimPrefix(line, prefix+":"))
		}
	}

	email.Body = strings.TrimSpace(strings.Join(bodyLines, "\n"))

	return *email, nil
}

func getHeader(key string, message *mail.Message) string {
	return message.Header.Get(key)
}

func handleError(path string, err error) (Email, error) {
	return Email{}, fmt.Errorf("error reading file (%s): %v", path, err)
}

func isEmpty(text string) bool {
	return text == " " || text == "\t"
}

func setHeaderValue(i *int, line *string, fileLines *[]string) {
	for {
		lines := *fileLines

		if *i >= len(lines) {
			break
		}

		nextLine := lines[*i+1]
		var nextLineFirstChar string

		if len(nextLine) > 0 {
			nextLineFirstChar = nextLine[0:1]
		}

		if !isEmpty(nextLineFirstChar) {
			break
		}

		*line = strings.TrimSpace(*line) + " " + strings.TrimSpace(nextLine)
		*i++
	}
}
