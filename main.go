package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("error: arg missing")
		os.Exit(1)
	}

	findExecutable := execNameWithExt(os.Args[1])

	paths := strings.Split(os.Getenv("PATH"), ";")
	if len(paths) == 0 {
		fmt.Println("error: failed to read PATH")
		os.Exit(1)
	}

	for _, p := range paths {
		p = strings.TrimSpace(p)
		findBin := filepath.Join(p, findExecutable)
		if exists(findBin) {
			fmt.Println(findBin)
			os.Exit(0)
		}
	}

	os.Exit(1)
}

func execNameWithExt(a string) string {

	ext := ".exe"
	last4 := ""
	if len(a) > len(ext) {
		last4 = strings.ToLower(a[len(a)-len(ext):])
	}
	if last4 != ext {
		a = a + ext
	}
	return a
}

func exists(path string) bool {

	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}
