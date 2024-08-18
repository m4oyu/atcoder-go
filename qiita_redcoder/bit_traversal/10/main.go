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

var a []int = make([]int, 20)
var q []int = make([]int, 200)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < m; i++ {
		scanner.Scan()
		q[i], _ = strconv.Atoi(scanner.Text())
	}

	for i := 0; i < m; i++ {
		found := false
		for j := 0; j < (1 << uint(n)); j++ {
			sum := 0
			for k := 0; k < n; k++ {
				if j&(1<<uint(k)) != 0 {
					sum += a[k]
				}
			}
			if q[i] == sum {
				found = true
				break
			}
		}
		if found {
			fmt.Fprintln(writer, "yes")
		} else {
			fmt.Fprintln(writer, "no")
		}
	}
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

func calculateArea(x1, y1, x2, y2, x3, y3, x4, y4 int) int {
	// Shoelace formula calculation
	area := math.Abs(float64(
		(x1*y2 + x2*y3 + x3*y4 + x4*y1) -
			(y1*x2 + y2*x3 + y3*x4 + y4*x1),
	))

	return int(area / 2)
}
