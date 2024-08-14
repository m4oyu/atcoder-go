package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
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

// 1 <= a, b, ab <= 5000
// 5000^3
// (a+b)/2 <= ab ならABピザを買う

func main() {
	scanner.Scan()
	a, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	b, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	c, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	y, _ := strconv.Atoi(scanner.Text())

	ans := math.MaxInt
	if (a+b)/2 <= c { // dont buy ab pizza
		ans = a*x + b*y
	} else {
		if x < y {
			for i := x * 2; i <= y*2; i++ {

				ans = min(ans, c*i+b*(y-(i/2)))
			}
		} else {
			for i := y * 2; i <= x*2; i++ {
				ans = min(ans, c*i+a*(x-(i/2)))
			}
		}
	}

	if ans < 0 {
		time.Sleep(2 * time.Second)
	}

	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
