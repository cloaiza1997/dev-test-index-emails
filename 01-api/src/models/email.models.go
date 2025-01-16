package models

type Email struct {
	MessageID string `json:"messageId"`
	Date      string `json:"date"`
	From      string `json:"from"`
	To        string `json:"to"`
	Cc        string `json:"cc"`
	Bcc       string `json:"bcc"`
	Subject   string `json:"subject"`
	XFrom     string `json:"xFrom"`
	XTo       string `json:"xTo"`
	XCc       string `json:"xCc"`
	XBcc      string `json:"xBcc"`
	XFolder   string `json:"xFolder"`
	XOrigin   string `json:"xOrigin"`
	XFileName string `json:"xFileName"`
	Body      string `json:"body"`
	Path      string `json:"path"`
}

type EmailList struct {
	Pagination Pagination `json:"pagination"`
	Items      []Email    `json:"items"`
}
