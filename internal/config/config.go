package config

import "time"

type Config struct {
	Request RequestConfig
	Count   int64
}

type RequestConfig struct {
	Url     string
	Method  string
	Header  []string
	Body    string
	Timeout time.Duration
}

func New(url string, method string, header []string, body string, count int64, timeout time.Duration) *Config {
	return &Config{
		Request: RequestConfig{
			Url:     url,
			Method:  method,
			Header:  header,
			Body:    body,
			Timeout: timeout,
		},
		Count: count,
	}
}
