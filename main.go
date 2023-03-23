package main

import (
	"fmt"
	"github.com/erfanmomeniii/ptest/cmd"
	_ "go.uber.org/automaxprocs"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	}
}
