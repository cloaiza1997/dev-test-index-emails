package functions

import (
	"fmt"
	"os"
	"strings"
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
