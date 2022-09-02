package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	pattern, str, err := readInput()
	if err != nil {
		fail(err)
	}
	isMatch := match(pattern, str)
	if !isMatch {
		os.Exit(0)
	}
	words := strings.Split(str, " ")
	fmt.Println(len(words))
}

func readInput() (pattern, src string, err error) {
	flag.StringVar(&pattern, "p", "", "pattern to match against")
	flag.Parse()
	if pattern == "" {
		return pattern, src, errors.New("missing pattern")
	}
	src = strings.Join(flag.Args(), "")
	if src == "" {
		return pattern, src, errors.New("missing string to match")
	}
	return pattern, src, nil
}

func fail(err error) {
	fmt.Println("match:", err)
	os.Exit(1)
}

func match(pattern string, src string) bool {
	return strings.Contains(src, pattern)
}
