package app

import (
	"bytes"
	"fmt"
	"github.com/ErfanMomeniii/colorful"
	"github.com/enescakir/emoji"
	"github.com/erfanmomeniii/ptest/internal/config"
	"github.com/erfanmomeniii/ptest/internal/util"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var (
	Reports        = make(chan Report)
	DiagramReports []time.Duration
)

type App struct {
	Config *config.Config
}

type Report struct {
	T              time.Duration
	IsSuccess      bool
	ResponseStatus int
}

func New(url string, method string, header []string, body string, timeout int64, count int64, chart bool) *App {
	return &App{
		Config: config.New(
			url, method, header, body, time.Duration(timeout), count, chart,
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
				T:              f.Sub(s),
				IsSuccess:      err == nil,
				ResponseStatus: statusCode,
			}
			DiagramReports = append(DiagramReports, report.T)
			reports <- report

			w.Done()
		}(Reports, wg)
	}

	wg.Wait()

	if a.Config.Diagram {
		DrawChart(DiagramReports)
	}

	close(Reports)
}

func PrintReports(reports <-chan Report, wg *sync.WaitGroup) {
	fmt.Println("------------------   Report   ------------------")

	for {
		if r, ok := <-reports; ok {
			if r.IsSuccess {
				colorful.Printf(
					colorful.GreenColor, colorful.DefaultBackground,
					"%v Status Code %d %v Response Time : %v\n", emoji.CheckMark,
					r.ResponseStatus, emoji.Stopwatch, r.T,
				)
			} else {
				colorful.Printf(
					colorful.RedColor, colorful.DefaultBackground,
					"%v Status Code %d\n", emoji.CrossMark, r.ResponseStatus,
				)
			}

			wg.Done()
		} else {
			break
		}
	}
}
func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}

func DrawChart(reports []time.Duration) {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		line := charts.NewLine()

		line.SetGlobalOptions(
			charts.WithInitializationOpts(
				opts.Initialization{
					Theme:           types.ThemeInfographic,
					BackgroundColor: "#fff8cd",
				},
			),
			charts.WithTitleOpts(
				opts.Title{Title: "Reports Diagram"},
			),
			charts.WithYAxisOpts(
				opts.YAxis{
					Name: "Cost time(s)",
					SplitLine: &opts.SplitLine{
						Show: false,
					},
				},
			),
			charts.WithXAxisOpts(
				opts.XAxis{Name: "Number"},
			),
		)

		xAxis := util.GenerateXAxis(len(reports))

		line.SetXAxis(xAxis).
			AddSeries("reports", util.GenerateLineItems(DiagramReports))

		line.Render(w)
	})

	colorful.Printf(
		colorful.YellowColor, colorful.DefaultBackground,
		"%v See diagram in 127.0.0.1:8081\n", emoji.ChartIncreasing,
	)

	http.ListenAndServe(":8081", nil)
}
