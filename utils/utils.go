package utils

import (
	"log"
	"os"
	"path"
	"runtime"
	"strings"
)

func ReadFile(filePath string) string {

	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("Could not find Caller of util.ReadFile")
	}

	// parse directory with pathFromCaller (which could be relative to Directory)
	absolutePath := path.Join(path.Dir(filename), filePath)
	body, err := os.ReadFile(absolutePath)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	return strings.TrimRight(string(body), "\n")
}
