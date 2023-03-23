package app

import (
	"bytes"
	"fmt"
	"github.com/ErfanMomeniii/colorful"
	"github.com/ErfanMomeniii/ptest/internal/config"
	"github.com/enescakir/emoji"
	"net/http"
	"sync"
	"time"
)

type App struct {
	Config *config.Config
}

type Report struct {
	t              time.Duration
	isSuccess      bool
	responseStatus int
}

func New(url string, method string, count int64, timeout int64) *App {
	return &App{
		Config: config.New(
			url, method, count, time.Duration(timeout),
		),
	}
}

func (a *App) Run() {
	var (
		i       int64
		mu      = new(sync.Mutex)
		wg      = new(sync.WaitGroup)
		reports []Report
	)

	for i = 0; i < a.Config.PTest.Count; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			defer mu.Unlock()

			s := time.Now()

			requestBody := bytes.NewBuffer([]byte{})

			client := http.Client{Timeout: a.Config.PTest.Timeout}

			req, _ := http.NewRequest(a.Config.PTest.Method, a.Config.PTest.Url, requestBody)

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

			reports = append(reports, report)
			wg.Done()
		}()
	}

	wg.Wait()

	PrintReports(reports)
}

func PrintReports(reports []Report) {
	fmt.Println("------------------   Report   ------------------")

	for _, r := range reports {
		if r.isSuccess {
			colorful.Printf(
				colorful.GreenColor, colorful.DefaultBackground,
				"%v Status Code %d %v Time Response : %v\n", emoji.CheckMark, r.responseStatus, emoji.Stopwatch, r.t,
			)
		} else {
			colorful.Printf(
				colorful.RedColor, colorful.DefaultBackground,
				"%v Status Code %d\n", emoji.CrossMark, r.responseStatus,
			)
		}
	}
}
