package util

import (
	"fmt"
	"github.com/go-echarts/go-echarts/v2/opts"
	"strings"
	"time"
)

func GenerateHeader(headers []string) (hs map[string][]string) {
	for _, header := range headers {
		i := strings.Index(header, ":")
		if i != -1 {
			hs[header[:i]] = strings.Split(header[i+1:], ",")
		}
	}

	return hs
}

func GenerateXAxis(count int) (result []string) {
	for i := 1; i <= count; i++ {
		result = append(result, fmt.Sprintf("%d", i))
	}

	return result
}

func GenerateLineItems(reports []time.Duration) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < len(reports); i++ {
		items = append(items, opts.LineData{Value: reports[i].Seconds()})
	}

	return items
}
