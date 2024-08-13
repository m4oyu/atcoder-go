package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func init() {
	if len(os.Args) >= 2 {
		if os.Args[1] == "debug" {
			debug()
		}
	}
	const buf = 200100
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, buf), buf)
}

func debug() {
	inputFile, err := os.Open("./.in")
	if err != nil {
		fmt.Fprintln(os.Stderr, "There is no inputFile.")
		os.Exit(1)
	}

	scanner = bufio.NewScanner(inputFile)
}

func main() {

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	ans := 0
	for i := 1; i <= n; i += 2 {
		eight := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				eight++
			}
		}
		if eight == 8 {
			ans++
		}
	}

	fmt.Println(ans)
}
