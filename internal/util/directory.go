package util

import (
	"os"
	"strings"
)

func CurrentDirectory() string {
	wd, _ := os.Getwd()
	wdArr := strings.Split(wd, "/")
	return wdArr[len(wdArr)-1]
}
