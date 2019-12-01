package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count bytes")
	flag.Parse()

	switch {
	case *bytes:
		fmt.Println("bytes: ", count(os.Stdin, bufio.ScanBytes))
	case *lines:
		fmt.Println("lines: ", count(os.Stdin, bufio.ScanLines))
	default:
		fmt.Println("words: ", count(os.Stdin, bufio.ScanWords))
	}
}

func count(r io.Reader, splitFn bufio.SplitFunc) int {
	sc := bufio.NewScanner(r)
	sc.Split(splitFn)

	var wc int
	for sc.Scan() {
		wc++
	}
	return wc
}
