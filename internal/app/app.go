package app

import (
	"bytes"
	"fmt"
	"github.com/ErfanMomeniii/colorful"
	"github.com/enescakir/emoji"
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

	statusCode := 500
	if resp != nil {
		statusCode = resp.StatusCode
	}

	report := Report{
		t:              f.Sub(s),
		isSuccess:      err == nil,
		responseStatus: statusCode,
	}

	report.Print()
}

func (r *Report) Print() {
	fmt.Println("------------------   Report   ------------------")
	if r.isSuccess {
		colorful.Printf(
			colorful.GreenColor, colorful.DefaultBackground,
			"%v Status Code %d %v Time Response : %v", emoji.CheckMark, r.responseStatus, emoji.Stopwatch, r.t,
		)
	} else {
		colorful.Printf(
			colorful.RedColor, colorful.DefaultBackground,
			"%v Status Code %d", emoji.CrossMark, r.responseStatus,
		)
	}
}
