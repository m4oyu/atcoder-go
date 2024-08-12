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

	for {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		x, _ := strconv.Atoi(scanner.Text())

		if n == 0 && x == 0 {
			return
		}

		ans := 0
		for i := 1; i <= n; i++ {
			for j := i + 1; j <= n; j++ {
				for k := j + 1; k <= n; k++ {
					if i+j+k == x {
						ans++
					}
				}
			}
		}
		fmt.Println(ans)
	}
}
