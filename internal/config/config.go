package config

import "time"

type Config struct {
	PTest PTest
}

type PTest struct {
	Url     string
	Method  string
	Count   int64
	Timeout time.Duration
}

func New(url string, method string, count int64, timeout time.Duration) *Config {
	return &Config{
		PTest: PTest{
			Url:     url,
			Method:  method,
			Count:   count,
			Timeout: timeout,
		},
	}
}
