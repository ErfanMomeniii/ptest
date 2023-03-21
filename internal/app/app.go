package app

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

type App struct {
	Url     string
	Method  string
	Timeout time.Duration
}

type Report struct {
	t              time.Duration
	isSuccess      bool
	responseStatus int
}

func New(url string, method string) *App {
	return &App{
		Url:    url,
		Method: method,
	}
}

func (a *App) Run() {
	s := time.Now()

	requestBody := bytes.NewBuffer([]byte{})

	client := http.Client{Timeout: a.Timeout}

	req, _ := http.NewRequest(a.Method, a.Url, requestBody)

	resp, err := client.Do(req)

	f := time.Now()

	report := Report{
		t:              f.Sub(s),
		isSuccess:      err == nil,
		responseStatus: resp.StatusCode,
	}

	report.Print()
}

func (r *Report) Print() {
	fmt.Println("Report")
}
