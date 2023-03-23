package util

import "strings"

func GenerateHeader(headers []string) (hs map[string][]string) {
	for _, header := range headers {
		i := strings.Index(header, ":")
		if i != -1 {
			hs[header[:i]] = strings.Split(header[i+1:], ";")
		}
	}

	return hs
}
