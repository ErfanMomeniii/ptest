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

var (
	Reports = make(chan Report)
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
		i  int64
		mu = new(sync.Mutex)
		wg = new(sync.WaitGroup)
	)

	go PrintReports(Reports, wg)

	for i = 0; i < a.Config.Count; i++ {
		wg.Add(2)

		go func(reports chan<- Report, w *sync.WaitGroup) {
			mu.Lock()
			defer mu.Unlock()

			s := time.Now()
			requestBody := bytes.NewBuffer([]byte(a.Config.Request.Body))

			client := http.Client{Timeout: a.Config.Request.Timeout}

			req, _ := http.NewRequest(a.Config.Request.Method, a.Config.Request.Url, requestBody)
			req.Header = util.GenerateHeader(a.Config.Request.Header)
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

			reports <- report

			w.Done()
		}(Reports, wg)
	}

	wg.Wait()

	close(Reports)
}

func PrintReports(reports <-chan Report, wg *sync.WaitGroup) {
	fmt.Println("------------------   Report   ------------------")

	for {
		if r, ok := <-reports; ok {
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

			wg.Done()
		} else {
			break
		}
	}
}
