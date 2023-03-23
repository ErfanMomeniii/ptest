package config

import "time"

type Config struct {
	Url     string
	Method  string
	Header  []string
	Body    string
	Count   int64
	Timeout time.Duration
}

func New(url string, method string, header []string, body string, count int64, timeout time.Duration) *Config {
	return &Config{
		Url:     url,
		Method:  method,
		Header:  header,
		Body:    body,
		Count:   count,
		Timeout: timeout,
	}
}
