package app

import (
	"bytes"
	"fmt"
	"github.com/ErfanMomeniii/colorful"
	"github.com/enescakir/emoji"
	"github.com/erfanmomeniii/ptest/internal/config"
	"github.com/erfanmomeniii/ptest/internal/util"
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

func New(url string, method string, header []string, body string, count int64, timeout int64) *App {
	return &App{
		Config: config.New(
			url, method, header, body, count, time.Duration(timeout),
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

	for i = 0; i < a.Config.Count; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			defer mu.Unlock()

			s := time.Now()
			requestBody := bytes.NewBuffer([]byte(a.Config.Body))

			client := http.Client{Timeout: a.Config.Timeout}

			req, _ := http.NewRequest(a.Config.Method, a.Config.Url, requestBody)
			req.Header = util.GenerateHeader(a.Config.Header)
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
