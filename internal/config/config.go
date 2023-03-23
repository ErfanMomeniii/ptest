package config

import "time"

type Config struct {
	Request RequestConfig
	Count   int64
	Diagram bool
}

type RequestConfig struct {
	Url     string
	Method  string
	Header  []string
	Body    string
	Timeout time.Duration
}

func New(url string, method string, header []string, body string, timeout time.Duration, count int64, Diagram bool) *Config {
	return &Config{
		Request: RequestConfig{
			Url:     url,
			Method:  method,
			Header:  header,
			Body:    body,
			Timeout: timeout,
		},
		Count:   count,
		Diagram: Diagram,
	}
}
