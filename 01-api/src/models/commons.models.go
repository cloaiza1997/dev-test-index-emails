package models

type Pagination struct {
	Total int `json:"total"`
	Count int `json:"count"`
	Pages int `json:"pages"`
	Prev  int `json:"prev"`
	Next  int `json:"next"`
}

type QuerySearch struct {
	Term  string
	Limit int
	Page  int
}

type Request struct {
	Method string
	Url    string
	Body   interface{}
	Auth   RequestAuth
}

type RequestAuth struct {
	User string
	Pass string
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
