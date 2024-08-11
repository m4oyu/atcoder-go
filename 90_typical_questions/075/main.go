package main

import (
	"bufio"
	"fmt"
	"math"
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
	var n float64
	scanner.Scan()
	n, _ = strconv.ParseFloat(scanner.Text(), 64)

	var a float64 = 2
	count := 0
	for math.Pow(a, 2) <= n {
		if math.Mod(n, a) == 0 {
			n /= a
			count++
		} else {
			a++
		}
	}

	if n > 1 { // n が1より大きい場合、素数が残っている
		count++
	}

	ans := 0
	tmp := 1
	for tmp < count {
		tmp *= 2
		ans++
	}
	fmt.Println(ans)
}
